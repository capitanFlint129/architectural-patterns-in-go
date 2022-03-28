package command

import (
	"syscall"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
)

type forkCommand struct {
	command       []string
	inputChannel  <-chan string
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (f *forkCommand) Execute() {
	if len(f.command) == 0 {
		f.errorChannel <- errors.ErrorNotEnoughArguments
	} else {
		executable := f.command[1]
		params := f.command[2:]
		_, err := syscall.ForkExec(executable, params, nil)
		if err != nil {
			f.errorChannel <- err
		}
	}
}

func NewForkCommand(
	command []string,
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &forkCommand{
		command:       command,
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
