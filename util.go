package esme

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"os"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func getRouteConfig(paths []string) ([]*config, error) {
	var configs []*config

	for _, path := range paths {
		fileBytes, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}

		var c config
		err = json.Unmarshal(fileBytes, &c)
		if err != nil {
			panic(err)
		}

		// extract env variables
		for _, g := range c.RouteGroups {
			// basic auth
			password, err := extractValueOrReturnInput(g.Auth.Basic.Password)
			if err != nil {
				panic(err)
			}
			g.Auth.Basic.Password = password

			username, err := extractValueOrReturnInput(g.Auth.Basic.Username)
			if err != nil {
				panic(err)
			}
			g.Auth.Basic.Username = username

			// bearer token
			token, err := extractValueOrReturnInput(g.Auth.BearerToken.Token)
			if err != nil {
				panic(err)
			}
			g.Auth.BearerToken.Token = token

			// custom
			for k, v := range g.Auth.Custom {
				value, err := extractValueOrReturnInput(v)
				if err != nil {
					panic(err)
				}

				g.Auth.Custom[k] = value
			}
		}

		err = verify(&c, path)
		if err != nil {
			panic(err)
		}

		configs = append(configs, &c)
	}

	return configs, nil
}

func extractValueOrReturnInput(input string) (string, error) {
	re, err := regexp.Compile(`\$\{(.+?)}`)
	if err != nil {
		slog.Error("Error compiling regex", "error", err)
		return input, err
	}

	match := re.FindStringSubmatch(input)
	if match != nil && len(match) > 1 {
		return match[1], nil
	}

	return input, nil
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

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return "Invalid validation error or other error type"
	}

	for _, err := range validationErrors {
		buffer.WriteString(err.Namespace() + ":" + err.ActualTag() + ";")
	}

	return buffer.String()
}
