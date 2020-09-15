package datastore

import "errors"

type Redis struct {
	Connection string
}

func (r Redis) GetFeeds() ([]Feed, error) {
	return nil, errors.New("Error occurred.")
}
