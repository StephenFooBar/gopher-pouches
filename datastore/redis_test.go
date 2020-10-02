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

func TestGetFeedsReturnsErrorWhenActiveFeedsKeyDoesNotExist(t *testing.T) {
	redis := GetInstance("host=:6379,database=" + NotInUseDB)
	expected := "redigo: nil returned"
	actual, err := redis.GetFeeds()
	assert.Nil(t, actual)
	assertError(t, err, expected)
}

func TestGetFeedsReturnsEmptyFeedWhenActiveFeedsKeyDoesExist(t *testing.T) {
	redis := GetInstance("host=:6379,database=" + TestDB)
	redis.CreateFeeds()
	actual, err := redis.GetFeeds()
	fmt.Println(err.Error())
	assert.NotNil(t, actual)
}

func assertError(t *testing.T, err error, expected string) {
	if assert.NotNil(t, err) {
		assert.Equal(t, expected, err.Error())
	}
}
