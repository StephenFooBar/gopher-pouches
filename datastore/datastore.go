package datastore

type Feed struct {
	Name string
	URL  string
}

type Datastore interface {
	GetFeeds() ([]Feed, error)
}
