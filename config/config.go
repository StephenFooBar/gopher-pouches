package config

import (
	"errors"
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

	if msg := validateConfig(*config); msg != command.Successful {
		return command.Response{msg, false}
	}
	return command.Response{command.Successful, true}
}

func validateConfig(config Config) string {
	if config.Datastore == "" {
		return command.ConfigEntryMissing
	}
	return command.Successful
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
