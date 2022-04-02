package command

import (
	"strings"
	"sync"
)

type echoCommand struct {
	args          []string
	inputChannel  <-chan string
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (e *echoCommand) Execute(wg *sync.WaitGroup) {
	defer wg.Done()
	e.outputChannel <- strings.Join(e.args, " ")
}

func NewEchoCommand(
	args []string,
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &echoCommand{
		args:          args,
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
