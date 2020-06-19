package main

type Response struct {
	message string
	success bool
}

type Command struct {
	command string
}

func main() {

}

func isValid(command Command) bool {
	return command.command == "list"
}

func RunCommand(command Command) Response {
	if !isValid(command) {
		return Response{"Invalid Command.", false}
	}
	return Response{"Successfully executed.", true}
}
