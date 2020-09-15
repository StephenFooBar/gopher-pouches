package datastore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFeedsReturnsErrorWhenConnectionIsEmpty(t *testing.T) {
	redis := Redis{""}
	actual, err := redis.GetFeeds()
	assert.Nil(t, actual)
	assert.NotNil(t, err)
}
