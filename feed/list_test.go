package feed

import (
	"testing"

	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/config"
	"github.com/stretchr/testify/assert"
)

func TestListFeedShouldShowErrorMessageWhenDataStoreIsNotSet(t *testing.T) {
	expected := command.DataStoreNotSet
	actual := List(config.Config{"", ""})
	assert.Equal(t, expected, actual.Message)
	assert.Equal(t, false, actual.Success)
}

func TestListFeedShouldShowErrorMessageWhenDataStoreNotSupported(t *testing.T) {
	expected := command.DataStoreNotSupported
	actual := List(config.Config{"not-existing-db", "not-existing-connection"})
	assert.Equal(t, expected, actual.Message)
	assert.Equal(t, false, actual.Success)
}

func TestListFeedShouldShowErrorMessageWhenErrorOccurredInFetchingFeedFromDataStore(t *testing.T) {
	if testing.Short() {
		t.Skip("")
	}
	expected := command.ErrorInDataStoreOperation
	actual := List(config.Config{"redis", ":6379"})
	assert.Equal(t, expected, actual.Message)
	assert.Equal(t, false, actual.Success)
}
