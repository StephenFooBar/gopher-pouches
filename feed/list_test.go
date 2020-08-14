package feed

import (
	"testing"

	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/config"
	"github.com/stretchr/testify/assert"
)

func TestListFeedShouldShowErrorMessageWhenDataStoreIsNotSet(t *testing.T) {
	expected := command.DataStoreNotSet
	actual := List(config.Config{""})
	assert.Equal(t, expected, actual.Message)
	assert.Equal(t, false, actual.Success)
}
