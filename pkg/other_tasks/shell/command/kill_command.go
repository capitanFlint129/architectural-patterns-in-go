package command

import (
	"context"
	errorTypes "github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
	"sync"
)

type killCommand struct {
	args          []string
	inputChannel  <-chan string
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (k *killCommand) Execute(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	//pid, err := strconv.Atoi(k.args[0])
	//if err != nil {
	//	k.errorChannel <- err
	//	return
	//}
	//if err = kill(pid); err != nil {
	//	k.errorChannel <- err
	//}
}

func (k *killCommand) SetArgs(args []string) error {
	switch len(k.args) {
	case 0:
		return errorTypes.ErrorNotEnoughArguments
	case 1:
		k.args = args
	default:
		return errorTypes.ErrorTooManyArguments
	}
	return nil
}

func NewKillCommand(
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &killCommand{
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
