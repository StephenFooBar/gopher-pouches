package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCommandShouldShowInvalidMessageWhenNilCommandIssued(t *testing.T) {
	expected := "Invalid Command."
	actual := RunCommand(Command{""})
	assert.Equal(t, expected, actual.message)
	if actual.message != "Invalid Command." {
		t.Errorf("Invalid output")
	}
}
