package datastore

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	NotInUseDB string = "1"
	TestDB     string = "2"
)

func TestGetFeedsReturnsErrorWhenConnectionIsEmpty(t *testing.T) {
	redis := GetInstance("")
	expected := EmptyConnection
	actual, err := redis.GetFeeds()
	assert.Nil(t, actual)
	assertError(t, err, expected)
}

func TestGetFeedsReturnsEmptyFeedWhenActiveFeedsKeyDoesNotExist(t *testing.T) {
	redis := GetInstance("host=:6379,database=" + NotInUseDB)
	actual, _ := redis.GetFeeds()
	assert.NotNil(t, actual)
	assert.Empty(t, actual)
}

func TestGetFeedsReturnsEmptyFeedWhenActiveFeedsKeyDoesExist(t *testing.T) {
	//t.Skip("skipping the test")
	redis := GetInstance("host=:6379,database=" + TestDB)
	redis.AddFeed("test")
	redis.RemoveFeed("test")
	actual, err := redis.GetFeeds()
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.NotNil(t, actual)
	assert.Empty(t, actual)
}

func assertError(t *testing.T, err error, expected string) {
	if assert.NotNil(t, err) {
		assert.Equal(t, expected, err.Error())
	}
}
