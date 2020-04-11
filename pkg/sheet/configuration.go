package sheet

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func ParseConfig() (Configuration, error) {
	var config Configuration

	configPath := os.Getenv("CONFIG_PATH")
	if len(configPath) < 1 {
		return config, fmt.Errorf("CONFIG_PATH must be set.")
	}

	config_data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config, fmt.Errorf("Could not parse config file at CONFIG_PATH.")
	}

	err = yaml.Unmarshal(config_data, &config)
	if err != nil {
		return config, fmt.Errorf("Failed to parse config file at CONFIG_PATH.")
	}

	return config, nil
}
