package episode

import (
	"testing"

	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/datastore/common"
	"github.com/StephenFooBar/gopher-pouches/test"
)

func TestListEpisodeShouldShowErrorWhenFeedIsNil(t *testing.T) {
	expected := command.MissingFeedInformation
	actual := List(common.Feed{})
	test.AssertFailure(t, expected, actual)
}
