package feed

import (
	"testing"

	"github.com/StephenFooBar/gopher-pouches/config"
	"github.com/stretchr/testify/assert"
)

func TestListFeedShouldShowErrorMessageWhenDataStoreIsNotSet(t *testing.T) {
	expected := "Data store is not set. Please run config to set data store information."
	actual := List(config.Config{""})
	assert.Equal(t, expected, actual.message)
	assert.Equal(t, false, actual.success)
}
