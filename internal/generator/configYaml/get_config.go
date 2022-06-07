package configYaml

import (
	"fmt"
	"github.com/terryhay/argtools/pkg/argtoolsError"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// GetConfig - loads a config yaml file and unmarshal it into Config object
func GetConfig(configPath string) (*Config, *argtoolsError.Error) {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, argtoolsError.NewError(
			argtoolsError.CodeGetConfigReadFileError,
			fmt.Errorf("configYaml.GetConfig: read config file error: %v", err))
	}
	config := new(Config)
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, argtoolsError.NewError(
			argtoolsError.CodeGetConfigUnmarshalError,
			fmt.Errorf("configYaml.GetConfig: unmarshal error: %v", err))
	}

	return config, nil
}
