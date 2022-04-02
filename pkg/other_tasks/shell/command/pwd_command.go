package command

import (
	"sync"
)

type pwdCommand struct {
	args          []string
	inputChannel  <-chan string
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (p *pwdCommand) Execute(wg *sync.WaitGroup) {
	defer wg.Done()
	wd, err := getwd()
	if err != nil {
		p.errorChannel <- err
	} else {
		p.outputChannel <- wd
	}
}

func NewPwdCommand(
	args []string,
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &pwdCommand{
		args:          args,
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
