package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/StephenFooBar/gopher-pouches/command"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Datastore string
}

const defaultYaml string = "config.yml"

func Get(configFilePath string) command.Response {
	if validateFile(configFilePath) != nil {
		return command.Response{command.ConfigFileMissing, false}
	}
	config, err := createConfig(configFilePath)
	if err != nil {
		return command.Response{command.InvalidConfig, false}
	}
	fmt.Println(config.Datastore)
	if config.Datastore == "" {
		return command.Response{command.DatastoreTypeMissing, false}
	}
	return command.Response{command.Successful, true}
}

func createConfig(filePath string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}

func validateFile(filePath string) error {
	s, err := os.Stat(filePath)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return errors.New("Given file path is a directory.")
	}
	return nil
}
