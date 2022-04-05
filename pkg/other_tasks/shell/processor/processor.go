package processor

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
	"sync"
)

type parser = interface {
	Parse(command string) []string
}

type command = interface {
	Execute(ctx context.Context, wg *sync.WaitGroup)
	SetArgs(args []string) error
}

type commandCreator = func(
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) command

type Processor interface {
	StartProcessing(ctx context.Context, wg *sync.WaitGroup)
}

type processor struct {
	pipeParser         parser
	commandParser      parser
	inputChannel       <-chan string
	outputChannel      chan<- string
	errorChannel       chan<- error
	commandCreatorsMap map[string]commandCreator
}

func (p *processor) StartProcessing(ctx context.Context, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for {
			select {
			case pipeCommand := <-p.inputChannel:
				parsedPipe := p.pipeParser.Parse(pipeCommand)
				preparedCommands, err := p.getPreparedCommands(parsedPipe)
				if err != nil {
					p.errorChannel <- err
					break
				}
				var wg sync.WaitGroup
				for _, preparedCommand := range preparedCommands {
					wg.Add(1)
					go preparedCommand.Execute(ctx, &wg)
				}
				wg.Wait()
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (p *processor) getPreparedCommands(parsedPipe []string) ([]command, error) {
	preparedCommands := make([]command, len(parsedPipe))

	if len(parsedPipe) == 1 {
		var err error
		preparedCommands[0], err = p.createCommand(parsedPipe[0], p.inputChannel, p.outputChannel, p.errorChannel)
		if err != nil {
			return nil, err
		}
		return preparedCommands, nil
	}

	var inputCommandChannel chan string
	var outputCommandChannel chan string

	for i, commandString := range parsedPipe {
		var err error
		switch i {
		case len(parsedPipe) - 1:
			preparedCommands[i], err = p.createCommand(commandString, inputCommandChannel, p.outputChannel, p.errorChannel)
		case 0:
			outputCommandChannel = make(chan string)
			preparedCommands[i], err = p.createCommand(commandString, p.inputChannel, outputCommandChannel, p.errorChannel)
		default:
			outputCommandChannel = make(chan string)
			preparedCommands[i], err = p.createCommand(commandString, inputCommandChannel, outputCommandChannel, p.errorChannel)
		}
		if err != nil {
			return nil, err
		}
		inputCommandChannel = outputCommandChannel
	}
	return preparedCommands, nil
}

func (p *processor) createCommand(commandString string, inputChannel <-chan string, outputChannel chan<- string, errorChannel chan<- error) (command, error) {
	parsedCommand := p.commandParser.Parse(commandString)
	if !p.checkCommandInMap(parsedCommand[0]) {
		return nil, errors.ErrorCommandNotFound
	}
	preparedCommand := p.commandCreatorsMap[parsedCommand[0]](inputChannel, outputChannel, errorChannel)
	err := preparedCommand.SetArgs(parsedCommand[1:])
	if err != nil {
		return nil, err
	}
	return preparedCommand, nil
}

func (p *processor) checkCommandInMap(command string) bool {
	if _, ok := p.commandCreatorsMap[command]; !ok {
		return false
	}
	return true
}

func NewProcessor(
	pipeParser parser,
	commandParser parser,
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
	commandCreatorsMap map[string]commandCreator,
) Processor {
	return &processor{
		pipeParser:         pipeParser,
		commandParser:      commandParser,
		inputChannel:       inputChannel,
		outputChannel:      outputChannel,
		errorChannel:       errorChannel,
		commandCreatorsMap: commandCreatorsMap,
	}
}
