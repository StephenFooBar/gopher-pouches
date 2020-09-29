package datastore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	NotInUseDB string = "1"
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

func assertError(t *testing.T, err error, expected string) {
	if assert.NotNil(t, err) {
		assert.Equal(t, expected, err.Error())
	}
}
