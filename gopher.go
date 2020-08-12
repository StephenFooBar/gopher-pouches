package main

import (
	"github.com/StephenFooBar/gopher-pouches/command"
)

func main() {

}

func isValid(command command.Command) bool {
	return command.Command == "list"
}

func RunCommand(cmd command.Command) command.Response {
	if !isValid(cmd) {
		return command.Response{"Invalid Command.", false}
	}
	return command.Response{"Successfully executed.", true}
}
