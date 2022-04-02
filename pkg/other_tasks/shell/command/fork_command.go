package command

import (
	errorTypes "github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
	"sync"
)

type forkCommand struct {
	args          []string
	inputChannel  <-chan string
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (f *forkCommand) Execute(wg *sync.WaitGroup) {
	defer wg.Done()
	if len(f.args) == 0 {
		f.errorChannel <- errorTypes.ErrorNotEnoughArguments
	} else {
		executable := f.args[0]
		args := f.args[1:]
		_, err := fork(executable, args)
		if err != nil {
			f.errorChannel <- err
		}
	}
}

func NewForkCommand(
	args []string,
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &forkCommand{
		args:          args,
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
