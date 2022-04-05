package command

import (
	"context"
	errorTypes "github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
	"sync"
)

type cdCommand struct {
	args          []string
	inputChannel  <-chan string
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (c *cdCommand) Execute(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	if len(c.args) == 0 {
		return
	}
	if err := chdir(c.args[0]); err != nil {
		c.errorChannel <- err
	}
}

func (c *cdCommand) SetArgs(args []string) error {
	if len(args) > 1 {
		return errorTypes.ErrorTooManyArguments
	}
	c.args = args
	return nil
}

func NewCdCommand(
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &cdCommand{
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
