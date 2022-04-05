package command

import (
	"context"
	"strings"
	"sync"
)

type echoCommand struct {
	args          []string
	inputChannel  <-chan string
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (e *echoCommand) Execute(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	e.outputChannel <- strings.Join(e.args, " ")
}

func (e *echoCommand) SetArgs(args []string) error {
	e.args = args
	return nil
}

func NewEchoCommand(
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &echoCommand{
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
