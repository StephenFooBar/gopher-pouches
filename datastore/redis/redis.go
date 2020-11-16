package redis

import (
	"errors"
	"strconv"
	"strings"

	"github.com/StephenFooBar/gopher-pouches/datastore"
	"github.com/gomodule/redigo/redis"
	redisDb "github.com/gomodule/redigo/redis"
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

func (r Redis) InitializeDb() error {
	conn, err := r.connect()
	if err != nil {
		return err
	}
	defer conn.Close()
	conn.Do("FLUSHDB")
	return nil
}

func (r Redis) AddFeed(feed datastore.Feed) error {
	conn, err := r.connect()
	if err != nil {
		return err
	}
	defer conn.Close()
	conn.Do("LPUSH", "ActiveFeeds", serializeFeed(feed))
	return nil
}

func (r Redis) RemoveFeed(feed datastore.Feed) error {
	conn, err := r.connect()
	if err != nil {
		return err
	}
	defer conn.Close()
	conn.Do("LREM", "ActiveFeeds", 0, serializeFeed(feed))
	return nil
}

func (r Redis) GetFeeds() ([]datastore.Feed, error) {
	conn, err := r.connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	s, err := redisDb.Strings(conn.Do("LRANGE", "ActiveFeeds", 0, -1))
	if err != nil {
		return nil, err
	}
	return convertToFeeds(s), nil
}

func convertToFeeds(s []string) []datastore.Feed {
	feeds := []datastore.Feed{}
	for _, feed := range s {
		feeds = append(feeds, deserializeFeed(feed))
	}
	return feeds
}

func serializeFeed(f datastore.Feed) string {
	var sb strings.Builder
	sb.WriteString("Name:")
	sb.WriteString(f.Name)
	sb.WriteString("|")
	sb.WriteString("URL:")
	sb.WriteString(f.URL)
	//sb.WriteString("|")
	return sb.String()
}

func deserializeFeed(s string) datastore.Feed {
	attr := strings.Split(s, "|")
	return datastore.Feed{
		Name: strings.TrimPrefix(attr[0], "Name:"),
		URL:  strings.TrimPrefix(attr[1], "URL:"),
	}
}

func (r Redis) connect() (redis.Conn, error) {
	if r.Connection == "" {
		return nil, errors.New(EmptyConnection)
	}
	conn, err := redisDb.Dial("tcp", r.Host)
	if err != nil {
		return nil, err
	}
	if r.Database != "" {
		db, _ := strconv.Atoi(r.Database)
		conn.Do("SELECT", db)
	}
	return conn, nil
}
