package feed

import (
	"testing"

	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/config"
	"github.com/StephenFooBar/gopher-pouches/datastore/common"
	"github.com/StephenFooBar/gopher-pouches/test"
	"github.com/stretchr/testify/assert"
)

func TestRemoveFeedShouldShowErrorMessageWhenDataStoreIsNotSet(t *testing.T) {
	expected := command.DataStoreNotSet
	actual := Remove(config.Config{"", ""}, test.NonRssMockFeed)
	test.AssertFailure(t, expected, actual)
}

func TestRemoveFeedShouldShowErrorMessageWhenDataStoreNotSupported(t *testing.T) {
	expected := command.DataStoreNotSupported
	actual := Remove(config.Config{"not-existing-db", "not-existing-connection"}, test.NonRssMockFeed)
	test.AssertFailure(t, expected, actual)
}

func TestRemoveFeedShouldShowErrorMessageWhenErrorOccurredWhileRemovingFeedInDataStore(t *testing.T) {
	expected := command.ErrorInDataStoreOperation
	failingPort := ":0000"
	actual := Remove(config.Config{"redis", failingPort}, test.NonRssMockFeed)
	test.AssertFailure(t, expected, actual)
}

func TestRemoveFeedShouldShowErrorMessageWhenFeedIsEmpty(t *testing.T) {
	expected := command.MissingFeedInformation
	conf := config.Config{"redis", test.ValidRedisConnection}
	test.InitializeRedis(conf)
	actual := Remove(conf, test.EmptyFeed)
	test.AssertFailure(t, expected, actual)
}

func TestRemoveFeedShouldShowErrorMessageWhenFeedIsMissingName(t *testing.T) {
	expected := command.MissingFeedInformation
	conf := config.Config{"redis", test.ValidRedisConnection}
	test.InitializeRedis(conf)
	actual := Remove(conf, common.Feed{"", "https://www.google.com"})
	test.AssertFailure(t, expected, actual)
}

/*
func TestAddFeedShouldShowErrorMessageWhenFeedIsMissingURL(t *testing.T) {
	expected := command.MissingFeedInformation
	conf := config.Config{"redis", validRedisConnection}
	test.InitializeRedis(conf)
	actual := Add(conf, common.Feed{"Mock Feed", ""})
	test.AssertFailure(t, expected, actual)
}
*/
func TestRemoveFeedShouldRemoveValidFeed(t *testing.T) {
	conf := config.Config{"redis", test.ValidRedisConnection}
	test.InitializeRedis(conf)
	Add(conf, test.NonRssMockFeed)
	original := List(conf)
	assert.Equal(t, true, original.Success)
	assert.Equal(t, 1, len(original.Data.([]common.Feed)))
	Remove(conf, test.NonRssMockFeed)
	actual := List(conf)
	assert.Equal(t, true, actual.Success)
	assert.Equal(t, 0, len(actual.Data.([]common.Feed)))
}
