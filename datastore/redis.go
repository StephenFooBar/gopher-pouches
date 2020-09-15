package datastore

import (
	"errors"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	Connection string
}

func (r Redis) GetFeeds() ([]Feed, error) {
	if r.Connection == "" {
		return nil, errors.New(EmptyConnection)
	}
	conn, err := redis.Dial("tcp", r.Connection)
	if err != nil {
		return nil, err
	}
	fmt.Print(conn)
	return nil, nil

	//return nil, errors.New("Error occurred.")
}
