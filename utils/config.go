package utils

import (
	"fmt"
	"os"
	"watchman/schema"

	"gopkg.in/yaml.v3"
)

func ReadConfig() schema.ConfigType {
	var Config schema.ConfigType

	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		fmt.Println(err)
	}

	return Config
}
