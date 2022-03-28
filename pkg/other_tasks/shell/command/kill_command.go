package command

import (
	"strconv"
	"syscall"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
)

type killCommand struct {
	command       []string
	inputChannel  <-chan string
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (k *killCommand) Execute() {
	switch len(k.command) {
	case 1:
		k.errorChannel <- errors.ErrorNotEnoughArguments
	case 2:
		pid, err := strconv.Atoi(k.command[1])
		if err != nil {
			k.errorChannel <- err
		} else {
			err = syscall.Kill(pid, syscall.SIGKILL)
			if err != nil {
				k.errorChannel <- err
			}
		}
	default:
		k.errorChannel <- errors.ErrorTooManyArguments

	}
}

func NewKillCommand(
	command []string,
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &killCommand{
		command:       command,
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
