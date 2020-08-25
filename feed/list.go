package feed

import (
	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/config"
)

func List(conf config.Config) command.Response {
	if config.ValidateConfig(conf) != command.Successful {
		return command.Response{command.DataStoreNotSet, false}
	}
	return command.Response{command.Successful, true}
}
