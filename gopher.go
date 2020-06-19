package main

type Response struct {
	message string
}

type Command struct {
	command string
}

func main() {

}

func RunCommand(command Command) Response {
	return Response{"Invalid Command."}
}
