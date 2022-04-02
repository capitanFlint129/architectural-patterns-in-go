package command

import (
	"sync"

	errorTypes "github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
)

type cdCommand struct {
	args          []string
	inputChannel  <-chan string
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (c *cdCommand) Execute(wg *sync.WaitGroup) {
	defer wg.Done()
	var err error
	switch len(c.args) {
	case 0:
	case 1:
		err = chdir(c.args[0])
	default:
		err = errorTypes.ErrorTooManyArguments
	}
	if err != nil {
		c.errorChannel <- err
	}
}

func NewCdCommand(
	args []string,
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &cdCommand{
		args:          args,
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
