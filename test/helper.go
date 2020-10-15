package test

import (
	"testing"

	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/stretchr/testify/assert"
)

func AssertFailure(t *testing.T, expected string, actual command.Response) {
	assert.False(t, actual.Success)
	assert.Equal(t, expected, actual.Message)
}
