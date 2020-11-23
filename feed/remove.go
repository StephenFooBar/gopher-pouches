package feed

import (
	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/config"
	"github.com/StephenFooBar/gopher-pouches/datastore"
	"github.com/StephenFooBar/gopher-pouches/datastore/common"
)

func Remove(conf config.Config, feed common.Feed) command.Response {
	if !validate(feed) {
		return command.Response{command.MissingFeedInformation, false, nil}
	}
	if config.Validate(conf) != command.Successful {
		return command.Response{command.DataStoreNotSet, false, nil}
	}
	ds := datastore.GetDatastore(conf)
	if ds == nil {
		return command.Response{command.DataStoreNotSupported, false, nil}
	}
	err := ds.RemoveFeed(feed)
	if err != nil {
		return command.Response{command.ErrorInDataStoreOperation, false, nil}
	}
	return command.Response{"", false, nil}
}
