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

const (
	RedisConnectionPrefix string = "host=:6379,database="
)

var mockFeed = Feed{"test name", "test url"}

func TestGetFeedsReturnsErrorWhenConnectionIsEmpty(t *testing.T) {
	redis := GetInstance("")
	expectedError := EmptyConnection
	actual, err := redis.GetFeeds()
	assertError(t, err, expectedError, actual)
}

func TestGetFeedsReturnsEmptyFeedWhenActiveFeedsKeyDoesNotExist(t *testing.T) {
	redis := GetInstance(RedisConnectionPrefix + NotInUseDB)
	actual, err := redis.GetFeeds()
	assertEmpty(t, err, actual)
}

func TestGetFeedsReturnsEmptyFeedWhenNothingIsInActiveFeeds(t *testing.T) {
	redis := GetInstance(RedisConnectionPrefix + TestDB)
	redis.InitializeDb()
	redis.AddFeed(mockFeed)
	redis.RemoveFeed(mockFeed)
	actual, err := redis.GetFeeds()
	redis.InitializeDb()
	assertEmpty(t, err, actual)
}

func TestGetFeedsReturnsAFeedWhenAFeedIsAdded(t *testing.T) {
	redis := GetInstance(RedisConnectionPrefix + TestDB)
	redis.InitializeDb()
	expected := mockFeed
	redis.AddFeed(expected)
	actual, _ := redis.GetFeeds()
	redis.InitializeDb()
	if assert.Len(t, actual, 1) {
		assert.Equal(t, expected, actual[0])
	}
	//	redis.InitializeDb()
}

func assertEmpty(t *testing.T, err error, actual []Feed) {
	if !assert.Nil(t, err) {
		fmt.Println(err.Error())
	}
	assert.NotNil(t, actual)
	assert.Empty(t, actual)
}

func assertError(t *testing.T, err error, expected string, actual []Feed) {
	assert.Nil(t, actual)
	if assert.NotNil(t, err) {
		assert.Equal(t, expected, err.Error())
	}
}
