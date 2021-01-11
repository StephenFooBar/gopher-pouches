package episode

import (
	"testing"

	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/datastore/common"
	"github.com/StephenFooBar/gopher-pouches/test"
	"github.com/stretchr/testify/assert"
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

func TestListEpisodeShouldShowErrorWhenFeedIsNotAValidRssFeed(t *testing.T) {
	t.Skip("Need to put other tests first")
	expected := command.InvalidFeed
	actual := List(test.NonRssMockFeed)
	test.AssertFailure(t, expected, actual)
}

func TestParseFeedAsRssShouldShowErrorWhenFeedHasInvalidURL(t *testing.T) {
	expected := command.InvalidURL
	_, actual := ParseFeedAsRss(common.Feed{"mockFeed", "not url"})
	assert.Equal(t, expected, actual)
}
