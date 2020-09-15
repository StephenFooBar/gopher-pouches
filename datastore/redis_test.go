package datastore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFeedsReturnsErrorWhenConnectionIsEmpty(t *testing.T) {
	redis := Redis{""}
	expected := EmptyConnection
	actual, err := redis.GetFeeds()
	assert.Nil(t, actual)
	assert.NotNil(t, err)
	assert.Equal(t, expected, err.Error())
}

func TestGetFeedsReturnsErrorWhenActiveFeedsKeyDoesNotExist(t *testing.T) {
	redis := Redis{":6379"}
	expected := FeedsListDoNotExist
	actual, err := redis.GetFeeds()
	assert.Nil(t, actual)
	assert.NotNil(t, err)
	assert.Equal(t, expected, err.Error())
}
