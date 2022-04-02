package command

import (
	"strconv"
	"sync"
	"syscall"

	errorTypes "github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
)

type killCommand struct {
	args          []string
	inputChannel  <-chan string
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (k *killCommand) Execute(wg *sync.WaitGroup) {
	defer wg.Done()
	switch len(k.args) {
	case 0:
		k.errorChannel <- errorTypes.ErrorNotEnoughArguments
	case 1:
		pid, err := strconv.Atoi(k.args[0])
		if err != nil {
			k.errorChannel <- err
		} else {
			err = syscall.Kill(pid, syscall.SIGKILL)
			if err != nil {
				k.errorChannel <- err
			}
		}
	default:
		k.errorChannel <- errorTypes.ErrorTooManyArguments

	}
}

func NewKillCommand(
	args []string,
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &killCommand{
		args:          args,
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
