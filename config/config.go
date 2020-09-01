package config

import (
	"errors"
	"os"

	"github.com/StephenFooBar/gopher-pouches/command"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Datastore  string
	Connection string
}

const defaultYaml string = "config.yml"

func Get(configFilePath string) (*Config, command.Response) {
	if validateFile(configFilePath) != nil {
		return nil, command.Response{command.ConfigFileMissing, false, nil}
	}
	config, err := createConfig(configFilePath)
	if err != nil {
		return nil, command.Response{command.InvalidConfig, false, nil}
	}

	if msg := Validate(*config); msg != command.Successful {
		return nil, command.Response{msg, false, nil}
	}
	return config, command.Response{command.Successful, true, nil}
}

func Validate(config Config) string {
	if config.Datastore == "" {
		return command.ConfigEntryMissing
	}
	if config.Connection == "" {
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
