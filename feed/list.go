package feed

import (
	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/config"
	"github.com/StephenFooBar/gopher-pouches/datastore"
)

func List(conf config.Config) command.Response {
	if config.Validate(conf) != command.Successful {
		return command.Response{command.DataStoreNotSet, false}
	}
	ds := getDatastore(conf)
	if ds == nil {
		return command.Response{command.DataStoreNotSupported, false}
	}
	/*
		feeds, err := ds.GetFeeds(conf.Connection)

		if err != nil {
			return command.Response{command.ErrorInDataStoreOperation, false}
		}*/

	return command.Response{command.Successful, true}
}

func getDatastore(conf config.Config) datastore.Datastore {
	switch conf.Datastore {
	default:
		return nil
	}
}
