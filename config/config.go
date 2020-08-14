package config

import (
	"github.com/StephenFooBar/gopher-pouches/command"
)

type Config struct {
	Datastore string
}

func Get() command.Response {
	return command.Response{command.YamlFileMissing, false}
}
