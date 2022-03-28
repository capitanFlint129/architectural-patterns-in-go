package command

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
	"os"
)

type cdCommand struct {
	command       []string
	inputChannel  <-chan string
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (c *cdCommand) Execute() {
	var err error
	switch len(c.command) {
	case 1:
	case 2:
		err = os.Chdir(c.command[1])
	default:
		err = errors.ErrorTooManyArguments
	}
	if err != nil {
		c.errorChannel <- err
	}
}

func NewCdCommand(
	command []string,
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &cdCommand{
		command:       command,
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
