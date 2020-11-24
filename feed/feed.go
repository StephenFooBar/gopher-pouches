package feed

import "github.com/StephenFooBar/gopher-pouches/datastore/common"

func Validate(feed common.Feed) bool {
	return feed.Name != "" && feed.URL != ""
}
