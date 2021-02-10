package episode

import (
	"testing"

	"github.com/StephenFooBar/gopher-pouches/command"
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
	actual := List(test.MissingURLMockFeed)
	test.AssertFailure(t, expected, actual)
}

func TestListEpisodeShouldShowErrorWhenFeedIsNotAValidRssFeed(t *testing.T) {
	t.Skip("Need to put other tests first")
	//must run mock servers here
	expected := command.InvalidFeed
	actual := List(test.NonRssMockFeed)
	test.AssertFailure(t, expected, actual)
}

func TestParseFeedAsRssShouldShowErrorWhenFeedHasInvalidURL(t *testing.T) {
	expected := command.InvalidURL
	_, actual := ParseFeedAsRss(test.InvalidURLMockFeed)
	assert.Equal(t, expected, actual)
}

func TestParseFeedAsRssShouldShowErrorWhenHttpResponseIsNotOK(t *testing.T) {
	expected := command.FailedRetrievingRss
	_, actual := ParseFeedAsRss(test.NotFoundURLMockFeed)
	assert.Equal(t, expected, actual)
}

func TestParseFeedAsRssShouldShowErrorWhenFeedIsMissingRssTag(t *testing.T) {
	srv, httpServerDoneExit := test.RunMockHttpServer(test.MockPort)
	expected := command.InvalidFeed
	_, actual := ParseFeedAsRss(test.MissingRssTagFeed)
	assert.Equal(t, expected, actual)
	test.StopMockHttpServer(srv, httpServerDoneExit)
}
