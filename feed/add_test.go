package feed

import (
	"testing"

	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/config"
	"github.com/StephenFooBar/gopher-pouches/datastore/common"
	"github.com/StephenFooBar/gopher-pouches/test"
	"github.com/stretchr/testify/assert"
)

func TestAddFeedShouldShowErrorMessageWhenDataStoreIsNotSet(t *testing.T) {
	expected := command.DataStoreNotSet
	actual := Add(config.Config{"", ""}, test.NonRssMockFeed)
	test.AssertFailure(t, expected, actual)
}

func TestAddFeedShouldShowErrorMessageWhenDataStoreNotSupported(t *testing.T) {
	expected := command.DataStoreNotSupported
	actual := Add(config.Config{"not-existing-db", "not-existing-connection"}, test.NonRssMockFeed)
	test.AssertFailure(t, expected, actual)
}

func TestAddFeedShouldShowErrorMessageWhenErrorOccurredWhileAddingFeedInDataStore(t *testing.T) {
	expected := command.ErrorInDataStoreOperation
	failingPort := ":0000"
	actual := Add(config.Config{"redis", failingPort}, test.NonRssMockFeed)
	test.AssertFailure(t, expected, actual)
}

func TestAddFeedShouldShowErrorMessageWhenFeedIsEmpty(t *testing.T) {
	expected := command.MissingFeedInformation
	conf := config.Config{"redis", test.ValidRedisConnection}
	test.InitializeRedis(conf)
	actual := Add(conf, test.EmptyFeed)
	test.AssertFailure(t, expected, actual)
}

func TestAddFeedShouldShowErrorMessageWhenFeedIsMissingName(t *testing.T) {
	expected := command.MissingFeedInformation
	conf := config.Config{"redis", test.ValidRedisConnection}
	test.InitializeRedis(conf)
	actual := Add(conf, common.Feed{"", "https://www.google.com"})
	test.AssertFailure(t, expected, actual)
}

func TestAddFeedShouldShowErrorMessageWhenFeedIsMissingURL(t *testing.T) {
	expected := command.MissingFeedInformation
	conf := config.Config{"redis", test.ValidRedisConnection}
	test.InitializeRedis(conf)
	actual := Add(conf, common.Feed{"Mock Feed", ""})
	test.AssertFailure(t, expected, actual)
}

func TestAddFeedShouldDisplayFeedWhenFeedIsValid(t *testing.T) {
	conf := config.Config{"redis", test.ValidRedisConnection}
	test.InitializeRedis(conf)
	Add(conf, test.NonRssMockFeed)
	actual := List(conf)
	assert.Equal(t, true, actual.Success)
	assert.Equal(t, test.NonRssMockFeed, actual.Data.([]common.Feed)[0])
}
