package main

import (
	"testing"

	"github.com/StephenFooBar/gopher-pouches/command"
	"github.com/StephenFooBar/gopher-pouches/config"
	"github.com/StephenFooBar/gopher-pouches/feed"
	"github.com/stretchr/testify/assert"
)

func TestRunCommandShouldShowInvalidMessageWhenNilCommandIssued(t *testing.T) {
	expected := command.InvalidCommand
	actual := RunCommand(command.Command{""})
	assert.Equal(t, expected, actual.Message)
	assert.Equal(t, false, actual.Success)
}

func TestRunCommandShouldShowInvalidMessageWhenAnInvalidCommandIssued(t *testing.T) {
	expected := command.InvalidCommand
	actual := RunCommand(command.Command{"Invalidcommand"})
	assert.Equal(t, expected, actual.Message)
	assert.Equal(t, false, actual.Success)
}

func TestRunCommandShouldReturnSuccessfullyWhenAValidCommandIssued(t *testing.T) {
	actual := RunCommand(command.Command{"list"})
	assert.Equal(t, true, actual.Success)
}

func TestCanCallListFromFeedPackage(t *testing.T) {
	feed.List(config.Config{"", ""})
}
