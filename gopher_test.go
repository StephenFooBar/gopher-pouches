package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCommandShouldShowInvalidMessageWhenNilCommandIssued(t *testing.T) {
	expected := "Invalid Command."
	actual := RunCommand(Command{""})
	assert.Equal(t, expected, actual.message)
}

func TestRunCommandShouldShowInvalidMessageWhenAnInvalidCommandIssued(t *testing.T) {
	expected := "Invalid Command."
	actual := RunCommand(Command{"Invalidcommand"})
	assert.Equal(t, expected, actual.message)
}

func TestRunCommandShouldReturnSuccessfullyWhenAValidCommandIssued(t *testing.T) {
	actual := RunCommand(Command{"list"})
	assert.Equal(t, true, actual.success)
}
