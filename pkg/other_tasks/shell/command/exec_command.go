package command

import (
	errorTypes "github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
	"sync"
)

type execCommand struct {
	args          []string
	inputChannel  <-chan string
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (e *execCommand) Execute(wg *sync.WaitGroup) {
	defer wg.Done()
	if len(e.args) == 0 {
		e.errorChannel <- errorTypes.ErrorNotEnoughArguments
	} else {
		executable := e.args[0]
		args := e.args[1:]
		err := exec(executable, args)
		if err != nil {
			e.errorChannel <- err
		}
	}

}

func NewExecCommand(
	args []string,
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &execCommand{
		args:          args,
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
