package datastore

import "github.com/StephenFooBar/gopher-pouches/config"

type Feed struct {
	Name string
	URL  string
}

type Datastore interface {
	GetFeeds() ([]Feed, error)
	AddFeed(feed Feed) error
	RemoveFeed(feed Feed) error
}

const (
	EmptyConnection     string = "Connection string is empty."
	FeedsListDoNotExist string = "Feeds List Do Not Exist."
)

func GetDatastore(conf config.Config) Datastore {
	switch conf.Datastore {
	case "redis":
		return GetInstance(conf.Connection)
	default:
		return nil
	}
}
