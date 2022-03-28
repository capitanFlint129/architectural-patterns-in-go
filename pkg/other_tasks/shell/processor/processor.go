package processor

import (
	"context"
	"sync"
)

type parser interface {
	Parse(command string) []string
}

type command = interface {
	Execute()
}

// TODO вопрос: где расположить публичный тип (или он не должен быть публичным?)
type CommandCreator = func(
	command []string,
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
	commandCreatorsMap map[string]CommandCreator
}

func (p *processor) StartProcessing(ctx context.Context, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for {
			select {
			case pipeCommand := <-p.inputChannel:
				parsedPipe := p.pipeParser.Parse(pipeCommand)
				preparedCommands := make([]command, len(parsedPipe))
				if len(parsedPipe) == 1 {
					parsedCommand := p.commandParser.Parse(parsedPipe[0])
					preparedCommands[0] = p.commandCreatorsMap[parsedCommand[0]](parsedCommand, p.inputChannel, p.outputChannel, p.errorChannel)
				} else {
					var inputCommandChannel chan string
					var outputCommandChannel chan string

					for i, command := range parsedPipe {
						parsedCommand := p.commandParser.Parse(command)
						switch i {
						case len(parsedPipe) - 1:
							preparedCommands[i] = p.commandCreatorsMap[parsedCommand[0]](parsedCommand, inputCommandChannel, p.outputChannel, p.errorChannel)
						case 0:
							outputCommandChannel = make(chan string)
							preparedCommands[i] = p.commandCreatorsMap[parsedCommand[0]](parsedCommand, p.inputChannel, outputCommandChannel, p.errorChannel)
						default:
							outputCommandChannel = make(chan string)
							preparedCommands[i] = p.commandCreatorsMap[parsedCommand[0]](parsedCommand, inputCommandChannel, outputCommandChannel, p.errorChannel)
						}
						inputCommandChannel = outputCommandChannel
					}
				}
				for _, preparedCommand := range preparedCommands {
					preparedCommand.Execute()
				}
			case <-ctx.Done():
				break
			}
		}
	}()
}

func NewProcessor(
	pipeParser parser,
	commandParser parser,
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
	commandCreatorsMap map[string]CommandCreator,
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
