package datastore

import (
	"errors"
	"strings"

	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	Connection string
	Host       string
	Database   string
}

func GetInstance(connection string) Redis {
	//separate by comma "redis0:6380,redis1:6380,allowAdmin=true"
	//host=:6379,database=0
	conn := strings.Split(connection, ",")
	host := getHost(conn)
	db := getDb(conn)
	return Redis{connection, host, db}
}

func getDb(conn []string) string {
	if len(conn) == 1 {
		return ""
	}
	db := strings.Split(conn[1], "=")
	return db[len(db)-1]
}

func getHost(conn []string) string {
	host := strings.Split(conn[0], "=")
	return host[len(host)-1]
}

func (r Redis) GetFeeds() ([]Feed, error) {
	if r.Connection == "" {
		return nil, errors.New(EmptyConnection)
	}
	conn, err := redis.Dial("tcp", r.Host)
	if err != nil {
		return nil, err
	}
	if conn != nil {
		return nil, nil

	}
	return nil, nil

	//return nil, errors.New("Error occurred.")
}
