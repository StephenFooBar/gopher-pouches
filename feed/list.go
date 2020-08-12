package feed

import (
	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/config"
)

func List(conf config.Config) command.Response {
	if conf.Datastore == "" {
		return command.Response{"Data store is not set. Please run config to set data store information.", false}
	}
	return command.Response{"Successful", true}
}
