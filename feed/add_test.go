package feed

import (
	"testing"

	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/config"
	"github.com/StephenFooBar/gopher-pouches/test"
)

func TestAddFeedShouldShowErrorMessageWhenDataStoreIsNotSet(t *testing.T) {
	expected := command.DataStoreNotSet
	actual := Add(config.Config{"", ""})
	test.AssertFailure(t, expected, actual)
}

func TestAddFeedShouldShowErrorMessageWhenDataStoreNotSupported(t *testing.T) {
	expected := command.DataStoreNotSupported
	actual := List(config.Config{"not-existing-db", "not-existing-connection"})
	test.AssertFailure(t, expected, actual)
}
