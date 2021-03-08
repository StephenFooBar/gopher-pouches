package episode

import (
	"os"
	"testing"

	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/test"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	srv, httpServerDoneExit := test.RunMockHttpServer(test.MockPort)
	code := m.Run()
	test.StopMockHttpServer(srv, httpServerDoneExit)
	os.Exit(code)
}

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
	expected := command.InvalidFeed
	actual := List(test.MissingRssTagFeed)
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
	expected := command.InvalidFeed
	_, actual := ParseFeedAsRss(test.MissingRssTagFeed)
	assert.Equal(t, expected, actual)
}

func TestParseFeedAsRssShouldShowErrorWhenFeedIsMissingChannelTag(t *testing.T) {
	expected := command.InvalidFeed
	_, actual := ParseFeedAsRss(test.MissingChannelTagFeed)
	assert.Equal(t, expected, actual)
}

func TestListEpisodeShouldShowNoEpisodeWhenFeedIsMissingItemTag(t *testing.T) {
	actual := List(test.MissingItemTagFeed)
	assert.Empty(t, actual.Data.(Rss).Items)
}
