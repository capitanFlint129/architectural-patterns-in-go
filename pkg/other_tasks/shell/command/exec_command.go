package command

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
	"syscall"
)

type execCommand struct {
	command       []string
	inputChannel  <-chan string
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (e *execCommand) Execute() {
	if len(e.command) == 0 {
		e.errorChannel <- errors.ErrorNotEnoughArguments
	} else {
		executable := e.command[1]
		params := e.command[2:]
		err := syscall.Exec(executable, params, nil)
		if err != nil {
			e.errorChannel <- err
		}
	}

}

func NewExecCommand(
	command []string,
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &execCommand{
		command:       command,
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
