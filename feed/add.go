package feed

import (
	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/config"
	"github.com/StephenFooBar/gopher-pouches/datastore"
)

func Add(conf config.Config, feed datastore.Feed) command.Response {
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
	err := ds.AddFeed(feed)
	if err != nil {
		return command.Response{command.ErrorInDataStoreOperation, false, nil}
	}
	return command.Response{"", false, nil}
}

func validate(feed datastore.Feed) bool {
	return feed.Name != "" && feed.URL != ""
}
