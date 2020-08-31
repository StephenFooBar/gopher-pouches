package datastore

type Feed struct {
	Name string
	URL  string
}

type Datastore interface {
	GetFeeds(connection string) ([]Feed, error)
}
