package command

import (
	"context"
	"sync"
)

type pwdCommand struct {
	args          []string
	inputChannel  <-chan string
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (p *pwdCommand) Execute(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	wd, err := getwd()
	if err != nil {
		p.errorChannel <- err
		return
	}
	p.outputChannel <- wd
}

func (p *pwdCommand) SetArgs(args []string) error {
	p.args = args
	return nil
}

func NewPwdCommand(
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &pwdCommand{
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
