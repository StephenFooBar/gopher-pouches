package datastore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFeedsReturnsErrorWhenConnectionIsEmpty(t *testing.T) {
	redis := GetInstance("")
	expected := EmptyConnection
	actual, err := redis.GetFeeds()
	assert.Nil(t, actual)
	assert.NotNil(t, err)
	assert.Equal(t, expected, err.Error())
}

func TestGetFeedsReturnsErrorWhenActiveFeedsKeyDoesNotExist(t *testing.T) {
	/*
		redis := GetInstance("host=:6379,database=1")
		expected := FeedsListDoNotExist
		actual, err := redis.GetFeeds()
		assert.Nil(t, actual)
		if assert.NotNil(t, err) {
			assert.Equal(t, expected, err.Error())
		}*/
}