package main

import (
	"testing"
)

func TestRunCommandShouldShowInvalidMessageWhenNilCommandIssued(t *testing.T) {
	actual := RunCommand(Command{""})
	if actual.message != "Invalid Command." {
		t.Errorf("Invalid output")
	}
}
