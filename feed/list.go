package feed

import (
	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/config"
)

func List(conf config.Config) command.Response {
	if config.Validate(conf) != command.Successful {
		return command.Response{command.DataStoreNotSet, false}
	}
	return command.Response{command.Successful, true}
}
