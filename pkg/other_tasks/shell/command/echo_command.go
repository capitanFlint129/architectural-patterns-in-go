package command

import "strings"

type echoCommand struct {
	command       []string
	inputChannel  <-chan string
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (e *echoCommand) Execute() {
	e.outputChannel <- strings.Join(e.command[1:], " ")
}

func NewEchoCommand(
	command []string,
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &echoCommand{
		command:       command,
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
