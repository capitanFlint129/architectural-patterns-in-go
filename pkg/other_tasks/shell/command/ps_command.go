package command

import (
	"context"
	"fmt"
	"sync"

	errorTypes "github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
)

const processFormat = "%d - %s"

type psCommand struct {
	args          []string
	inputChannel  <-chan string
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (p *psCommand) Execute(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	processes, err := ps()
	if err != nil {
		p.errorChannel <- err
		return
	}
	for _, process := range processes {
		p.outputChannel <- fmt.Sprintf(processFormat, process.pid, process.executable)
	}
}

func (p *psCommand) SetArgs(args []string) error {
	if len(args) > 0 {
		return errorTypes.ErrorTooManyArguments
	}
	p.args = args
	return nil
}

func NewPsCommand(
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &psCommand{
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
