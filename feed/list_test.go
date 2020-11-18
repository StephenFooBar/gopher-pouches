package feed

import (
	"testing"

	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/config"
	"github.com/StephenFooBar/gopher-pouches/datastore/common"
	"github.com/StephenFooBar/gopher-pouches/test"
	"github.com/stretchr/testify/assert"
)

func TestListFeedShouldShowErrorMessageWhenDataStoreIsNotSet(t *testing.T) {
	expected := command.DataStoreNotSet
	actual := List(config.Config{"", ""})
	test.AssertFailure(t, expected, actual)
}

func TestListFeedShouldShowErrorMessageWhenDataStoreNotSupported(t *testing.T) {
	expected := command.DataStoreNotSupported
	actual := List(config.Config{"not-existing-db", "not-existing-connection"})
	test.AssertFailure(t, expected, actual)
}

func TestListFeedShouldShowErrorMessageWhenErrorOccurredInFetchingFeedFromDataStore(t *testing.T) {
	expected := command.ErrorInDataStoreOperation
	failingPort := ":0000"
	actual := List(config.Config{"redis", failingPort})
	test.AssertFailure(t, expected, actual)
}

func TestListFeedShowReturnEmptyFeedWhenNoFeedExists(t *testing.T) {
	redisConnection := "host=:6379,database=2"
	conf := config.Config{"redis", redisConnection}
	test.InitializeRedis(conf)
	actual := List(conf)
	assert.Equal(t, true, actual.Success)
	assert.Empty(t, actual.Data.([]common.Feed))
}
