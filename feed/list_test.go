package feed

import (
	"testing"

	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/config"
	"github.com/StephenFooBar/gopher-pouches/datastore"
	"github.com/stretchr/testify/assert"
)

func TestListFeedShouldShowErrorMessageWhenDataStoreIsNotSet(t *testing.T) {
	expected := command.DataStoreNotSet
	actual := List(config.Config{"", ""})
	assertFailure(t, expected, actual)
}

func TestListFeedShouldShowErrorMessageWhenDataStoreNotSupported(t *testing.T) {
	expected := command.DataStoreNotSupported
	actual := List(config.Config{"not-existing-db", "not-existing-connection"})
	assertFailure(t, expected, actual)
}

func TestListFeedShouldShowErrorMessageWhenErrorOccurredInFetchingFeedFromDataStore(t *testing.T) {
	expected := command.ErrorInDataStoreOperation
	failingPort := ":0000"
	actual := List(config.Config{"redis", failingPort})
	assertFailure(t, expected, actual)
}

func TestListFeedShowReturnEmptyFeedWhenNoFeedExists(t *testing.T) {
	redisConnection := "host=:6379,database=2"
	actual := List(config.Config{"redis", redisConnection})
	assert.Equal(t, true, actual.Success)
	assert.Empty(t, actual.Data.([]datastore.Feed))
}

func assertFailure(t *testing.T, expected string, actual command.Response) {
	assert.False(t, actual.Success)
	assert.Equal(t, expected, actual.Message)
}
