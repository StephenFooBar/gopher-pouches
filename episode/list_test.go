package episode

import (
	"testing"

	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/datastore/common"
	"github.com/StephenFooBar/gopher-pouches/test"
)

func TestListEpisodeShouldShowErrorWhenFeedIsEmpty(t *testing.T) {
	expected := command.MissingFeedInformation
	actual := List(test.EmptyFeed)
	test.AssertFailure(t, expected, actual)
}

func TestListEpisodeShouldShowErrorWhenFeedIsMissingURL(t *testing.T) {
	expected := command.MissingFeedInformation
	actual := List(common.Feed{"mockFeed", ""})
	test.AssertFailure(t, expected, actual)
}

/*
func TestListEpisodeShouldShowErrorWhenFeedIsNotAValidRssFeed(t *testing.T) {
	expected := command.InvalidFeed
	actual := List(common.)
}
*/
