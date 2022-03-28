package command

import "os"

type pwdCommand struct {
	command       []string
	inputChannel  <-chan string
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (p *pwdCommand) Execute() {
	path, err := os.Getwd()
	if err != nil {
		p.errorChannel <- err
	}
	p.outputChannel <- path
}

func NewPwdCommand(
	command []string,
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &pwdCommand{
		command:       command,
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
