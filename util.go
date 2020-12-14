package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func getRouteConfig(paths []string) []config {
	var configs []config

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

		configs = append(configs, c)
	}

	return configs
}
