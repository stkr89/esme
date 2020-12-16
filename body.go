package esme

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
)

const InvalidRequestBodyError = "invalid request body"

func checkBody(requiredBody map[string]string, body io.ReadCloser) error {
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	var data map[string]interface{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}

	for k, v := range requiredBody {
		if _, ok := data[k]; !ok {
			return errors.New(InvalidRequestBodyError)
		}

		val := data[k]
		if val != v {
			return errors.New(InvalidRequestBodyError)
		}

	}

	return nil
}