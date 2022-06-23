package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func getVariables(data *Model, variablesFile string) error {
	if variablesFile != "" {
		err := readVariables(data, variablesFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func addEnvVariables(data *Model, envPrefix string) {
	envs := envToMap()
	for k, v := range envs {
		if strings.HasPrefix(k, envPrefix) {
			data.Variables[k] = v
		}
	}
}

func readVariables(m *Model, variables string) error {
	yamlFile, err := ioutil.ReadFile(variables)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, m)
	if err != nil {
		return fmt.Errorf("yamlFile.Unmarshal error   #%w ", err)
	}

	return nil
}

func envToMap() map[string]string {
	envMap := make(map[string]string)

	for _, v := range os.Environ() {
		splitV := strings.SplitN(v, "=", 2)
		envMap[splitV[0]] = splitV[1]
	}

	return envMap
}
