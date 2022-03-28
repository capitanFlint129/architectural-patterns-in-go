package command

import (
	"fmt"
	"github.com/mitchellh/go-ps"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
)

type psCommand struct {
	command       []string
	inputChannel  <-chan string
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (p *psCommand) Execute() {
	if len(p.command) > 1 {
		p.errorChannel <- errors.ErrorTooManyArguments
	} else {
		processes, err := ps.Processes()
		if err != nil {
			p.errorChannel <- err
		}
		for _, process := range processes {
			p.outputChannel <- fmt.Sprintf("%d - %s", process.Pid(), process.Executable())
		}
	}
}

func NewPsCommand(
	command []string,
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &psCommand{
		command:       command,
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
