package episode

import (
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
