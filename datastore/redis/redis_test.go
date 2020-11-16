package redis

import (
	"fmt"
	"testing"

	"github.com/StephenFooBar/gopher-pouches/datastore/common"
	"github.com/stretchr/testify/assert"
)

const (
	NotInUseDB string = "1"
	TestDB     string = "2"
)

const (
	RedisConnectionPrefix string = "host=:6379,database="
)

var mockFeed = common.Feed{"test name", "test url"}

func TestGetFeedsReturnsErrorWhenConnectionIsEmpty(t *testing.T) {
	r := GetInstance("")
	expectedError := common.EmptyConnection
	actual, err := r.GetFeeds()
	assertError(t, err, expectedError, actual)
}

func TestGetFeedsReturnsEmptyFeedWhenActiveFeedsKeyDoesNotExist(t *testing.T) {
	r := GetInstance(RedisConnectionPrefix + NotInUseDB)
	actual, err := r.GetFeeds()
	assertEmpty(t, err, actual)
}

func TestGetFeedsReturnsEmptyFeedWhenNothingIsInActiveFeeds(t *testing.T) {
	r := GetInstance(RedisConnectionPrefix + TestDB)
	r.InitializeDb()
	r.AddFeed(mockFeed)
	r.RemoveFeed(mockFeed)
	actual, err := r.GetFeeds()
	r.InitializeDb()
	assertEmpty(t, err, actual)
}

func TestGetFeedsReturnsAFeedWhenAFeedIsAdded(t *testing.T) {
	r := GetInstance(RedisConnectionPrefix + TestDB)
	r.InitializeDb()
	expected := mockFeed
	r.AddFeed(expected)
	actual, _ := r.GetFeeds()
	r.InitializeDb()
	if assert.Len(t, actual, 1) {
		assert.Equal(t, expected, actual[0])
	}
}

func assertEmpty(t *testing.T, err error, actual []common.Feed) {
	if !assert.Nil(t, err) {
		fmt.Println(err.Error())
	}
	assert.NotNil(t, actual)
	assert.Empty(t, actual)
}

func assertError(t *testing.T, err error, expected string, actual []common.Feed) {
	assert.Nil(t, actual)
	if assert.NotNil(t, err) {
		assert.Equal(t, expected, err.Error())
	}
}
