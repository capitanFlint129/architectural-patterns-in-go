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
	Execute(wg *sync.WaitGroup)
}

type commandCreator = func(
	args []string,
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
		selectCase:
			select {
			case pipeCommand := <-p.inputChannel:
				parsedPipe := p.pipeParser.Parse(pipeCommand)
				preparedCommands := make([]command, len(parsedPipe))

				if len(parsedPipe) == 1 {
					parsedCommand := p.commandParser.Parse(parsedPipe[0])
					if p.checkCommandInMap(parsedCommand[0]) {
						preparedCommands[0] = p.commandCreatorsMap[parsedCommand[0]](parsedCommand[1:], p.inputChannel, p.outputChannel, p.errorChannel)
					} else {
						p.errorChannel <- errors.ErrorCommandNotFound
						break
					}
				} else {
					var inputCommandChannel chan string
					var outputCommandChannel chan string

					for i, command := range parsedPipe {
						parsedCommand := p.commandParser.Parse(command)
						if !p.checkCommandInMap(parsedCommand[0]) {
							p.errorChannel <- errors.ErrorCommandNotFound
							break selectCase
						}
						switch i {
						case len(parsedPipe) - 1:
							preparedCommands[i] = p.commandCreatorsMap[parsedCommand[0]](parsedCommand[1:], inputCommandChannel, p.outputChannel, p.errorChannel)
						case 0:
							outputCommandChannel = make(chan string)
							preparedCommands[i] = p.commandCreatorsMap[parsedCommand[0]](parsedCommand[1:], p.inputChannel, outputCommandChannel, p.errorChannel)
						default:
							outputCommandChannel = make(chan string)
							preparedCommands[i] = p.commandCreatorsMap[parsedCommand[0]](parsedCommand[1:], inputCommandChannel, outputCommandChannel, p.errorChannel)
						}
						inputCommandChannel = outputCommandChannel
					}
				}

				var wg sync.WaitGroup
				for _, preparedCommand := range preparedCommands {
					wg.Add(1)
					go preparedCommand.Execute(&wg)
				}
				wg.Wait()
			case <-ctx.Done():
				return
			}
		}
	}()
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
