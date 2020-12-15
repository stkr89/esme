package esme

import (
	"bytes"
	"errors"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func getRouteConfig(paths []string) ([]*config, error) {
	var configs []*config

	for _, path := range paths {
		fileBytes, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		var c config
		err = yaml.Unmarshal(fileBytes, &c)
		if err != nil {
			panic(err)
		}

		err = verify(&c, path)
		if err != nil {
			return nil, err
		}

		configs = append(configs, &c)
	}

	return configs, nil
}

func verify(config *config, path string) error {
	err := validator.New().Struct(config)
	if err != nil {
		errStr := formatValidationError(err)
		if errStr != "" {
			return errors.New(path + " : " + errStr)
		}
	}

	return nil
}

func formatValidationError(err error) string {
	var buffer bytes.Buffer

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return buffer.String()
	}

	for _, err := range err.(validator.ValidationErrors) {
		buffer.WriteString(err.Namespace() + ":" + err.ActualTag() + ";")
	}

	return buffer.String()
}
