package episode

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/datastore/common"
	"github.com/StephenFooBar/gopher-pouches/feed"
)

func List(f common.Feed) command.Response {
	if !feed.Validate(f) {
		return command.Response{command.MissingFeedInformation, false, nil}
	}

	return command.Response{"success", true, nil}
}

func ParseFeedAsRss(f common.Feed) (Rss, string) {
	var rss Rss
	resp, err := http.Get(f.URL)
	if err != nil {
		return rss, command.InvalidURL
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return rss, command.FailedRetrievingRss
	}
	bodyInBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return rss, command.InvalidFeed
	}
	fmt.Println(string(bodyInBytes))
	return rss, command.Successful
	//todo: above correctly parse the rss feed
}
