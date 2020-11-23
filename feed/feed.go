package feed

import "github.com/StephenFooBar/gopher-pouches/datastore/common"

func validate(feed common.Feed) bool {
	return feed.Name != "" && feed.URL != ""
}
