package feed

import (
	"testing"

	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/config"
	"github.com/StephenFooBar/gopher-pouches/datastore/common"
	"github.com/StephenFooBar/gopher-pouches/test"
)

var (
	emptyFeed = common.Feed{}
	mockFeed  = common.Feed{"Mock Feed", "https://www.google.com"}
)

func TestAddFeedShouldShowErrorMessageWhenDataStoreIsNotSet(t *testing.T) {
	expected := command.DataStoreNotSet
	actual := Add(config.Config{"", ""}, mockFeed)
	test.AssertFailure(t, expected, actual)
}

func TestAddFeedShouldShowErrorMessageWhenDataStoreNotSupported(t *testing.T) {
	expected := command.DataStoreNotSupported
	actual := Add(config.Config{"not-existing-db", "not-existing-connection"}, mockFeed)
	test.AssertFailure(t, expected, actual)
}

func TestAddFeedShouldShowErrorMessageWhenErrorOccurredWhileAddingFeedInDataStore(t *testing.T) {
	expected := command.ErrorInDataStoreOperation
	failingPort := ":0000"
	actual := Add(config.Config{"redis", failingPort}, mockFeed)
	test.AssertFailure(t, expected, actual)
}

func TestAddFeedShouldShowErrorMessageWhenFeedIsEmpty(t *testing.T) {
	expected := command.MissingFeedInformation
	redisConnection := "host=:6379,database=2"
	conf := config.Config{"redis", redisConnection}
	test.InitializeRedis(conf)
	actual := Add(conf, emptyFeed)
	test.AssertFailure(t, expected, actual)
}
