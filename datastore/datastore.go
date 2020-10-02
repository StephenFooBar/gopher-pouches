package datastore

type Feed struct {
	Name string
	URL  string
}

type Datastore interface {
	GetFeeds() ([]Feed, error)
	CreateFeeds()
}

const (
	EmptyConnection     string = "Connection string is empty."
	FeedsListDoNotExist string = "Feeds List Do Not Exist."
)
