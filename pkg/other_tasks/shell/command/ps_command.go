package command

import (
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

func (p *psCommand) Execute(wg *sync.WaitGroup) {
	defer wg.Done()
	if len(p.args) > 0 {
		p.errorChannel <- errorTypes.ErrorTooManyArguments
	} else {
		processes, err := ps()
		fmt.Printf("%d - %d", processes, err)
		if err != nil {
			p.errorChannel <- err
		} else {
			for _, process := range processes {
				p.outputChannel <- fmt.Sprintf(processFormat, process.pid, process.executable)
			}
		}
	}
}

func NewPsCommand(
	args []string,
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &psCommand{
		args:          args,
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
