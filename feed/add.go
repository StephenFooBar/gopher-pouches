package feed

import (
	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/config"
	"github.com/StephenFooBar/gopher-pouches/datastore"
)

func Add(conf config.Config) command.Response {
	if config.Validate(conf) != command.Successful {
		return command.Response{command.DataStoreNotSet, false, nil}
	}
	ds := datastore.GetDatastore(conf)
	if ds == nil {
		return command.Response{command.DataStoreNotSupported, false, nil}
	}
	err := ds.AddFeed(datastore.Feed{})
	if err != nil {
		return command.Response{command.ErrorInDataStoreOperation, false, nil}
	}
	return command.Response{"", false, nil}
}
