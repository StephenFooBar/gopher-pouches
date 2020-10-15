package feed

import (
	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/config"
	"github.com/StephenFooBar/gopher-pouches/datastore"
)

func List(conf config.Config) command.Response {
	if config.Validate(conf) != command.Successful {
		return command.Response{command.DataStoreNotSet, false, nil}
	}
	ds := datastore.GetDatastore(conf)
	if ds == nil {
		return command.Response{command.DataStoreNotSupported, false, nil}
	}
	feeds, err := ds.GetFeeds()
	if err != nil {
		return command.Response{command.ErrorInDataStoreOperation, false, err}
	}
	return command.Response{command.Successful, true, feeds}
}

/*
func getDatastore(conf config.Config) datastore.Datastore {
	switch conf.Datastore {
	case "redis":
		return datastore.GetInstance(conf.Connection)
	default:
		return nil
	}
}*/
