package processor

import (
	"context"
	"errors"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	errorTypes "github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/processor/mocks"
)

const (
	runOneCommandTestCaseName         = "Run one command"
	runPipeTestCaseName               = "Run pipe command"
	unknownCommandTestCaseName        = "Command not exists"
	unknownCommandInPipeTestCaseName  = "Command in pipe not exists"
	commandExecutionErrorTestCaseName = "Error occurs while command executes"
)

var commandError = errors.New("Command error")

type consoleInput struct {
	commandString     string
	commandInputLines []string
}

type parsedCommand = []string

type inputData struct {
	consoleInput          consoleInput
	parsedPipe            []string
	parsedCommands        []parsedCommand
	processorCommands     []string
	commandExecutionError error
}

type expectedResult struct {
	commandExecutionNumber int
	outputChannelData      []string
	errorChannelErrors     []error
}

type testCommand struct {
	args              []string
	inputChannel      <-chan string
	outputChannel     chan<- string
	errorChannel      chan<- error
	executionNumber   int
	executionTestFunc func(inputChannel <-chan string, outputChannel chan<- string, errorChannel chan<- error, executionNumber int)
}

func (t *testCommand) Execute(wg *sync.WaitGroup) {
	defer wg.Done()
	t.executionTestFunc(t.inputChannel, t.outputChannel, t.errorChannel, t.executionNumber)
}

// Test_Processor creates processor, starts command processing and checks that
// commands send each other correct data
func Test_Processor(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      inputData
		expectedResult expectedResult
	}{
		{
			testCaseName: runOneCommandTestCaseName,
			inputData: inputData{
				consoleInput: consoleInput{
					commandString:     "command1 arg1 arg2\n",
					commandInputLines: []string{"line1", "line2"},
				},
				parsedPipe: []string{"command1 arg1 arg2"},
				parsedCommands: []parsedCommand{
					{"command1", "arg1", "arg2"},
				},
				processorCommands: []string{"command1"},
			},
			expectedResult: expectedResult{
				commandExecutionNumber: 1,
				outputChannelData:      []string{"line1", "line2"},
			},
		},
		{
			testCaseName: runPipeTestCaseName,
			inputData: inputData{
				consoleInput: consoleInput{
					commandString:     "command1 arg1 arg2 | command2 | command3\n",
					commandInputLines: []string{"line1", "line2"},
				},
				parsedPipe: []string{"command1  arg1 arg2", "command2", "command3"},
				parsedCommands: []parsedCommand{
					{"command1", "arg1", "arg2"},
					{"command2"},
					{"command3"},
				},
				processorCommands: []string{"command1", "command2", "command3"},
			},
			expectedResult: expectedResult{
				commandExecutionNumber: 1,
				outputChannelData:      []string{"line1", "line2"},
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			// Prepare parsers mocks
			pipeParserMock := mocks.NewParserMock()
			commandParserMock := mocks.NewParserMock()
			pipeParserMock.On("Parse", mock.AnythingOfType("string")).Return(testData.inputData.parsedPipe)
			commandParserMock.On("Parse", mock.AnythingOfType("string")).Return(func(commandString string) []string {
				for _, parsedCommand := range testData.inputData.parsedCommands {
					if strings.HasPrefix(commandString, parsedCommand[0]) {
						return parsedCommand
					}
				}
				return []string{}
			})
			// Prepare channels
			inputChannel := make(chan string)
			outputChannel := make(chan string)
			errorChannel := make(chan error)
			// Prepare commandCreators
			commandCreatorsMap := make(map[string]commandCreator)
			commandCreator := func(
				args []string,
				commandInputChannel <-chan string,
				commandOutputChannel chan<- string,
				commandErrorChannel chan<- error,
			) command {
				return &testCommand{
					args:            args,
					inputChannel:    commandInputChannel,
					outputChannel:   commandOutputChannel,
					errorChannel:    commandErrorChannel,
					executionNumber: 0,
					executionTestFunc: func(
						commandInputChannel <-chan string,
						commandOutputChannel chan<- string,
						commandErrorChannel chan<- error,
						executionNumber int,
					) {
						executionNumber++
						assert.Equal(t, testData.expectedResult.commandExecutionNumber, executionNumber)
						for i := range testData.inputData.consoleInput.commandInputLines {
							data := <-commandInputChannel
							assert.Equal(t, testData.inputData.consoleInput.commandInputLines[i], data)
							commandOutputChannel <- data
						}
					},
				}
			}
			for _, processorCommand := range testData.inputData.processorCommands {
				commandCreatorsMap[processorCommand] = commandCreator
			}
			// Create processor
			processor := NewProcessor(
				pipeParserMock,
				commandParserMock,
				inputChannel,
				outputChannel,
				errorChannel,
				commandCreatorsMap,
			)
			mainCtx := context.Background()
			ctx, cancel := context.WithCancel(mainCtx)
			var wg sync.WaitGroup

			// Start processing and send command string
			wg.Add(1)
			processor.StartProcessing(ctx, &wg)
			inputChannel <- testData.inputData.consoleInput.commandString

			// Write data to start of pipe and read it at the end of pipe
			resultData := make([]string, 0)
			for i := range testData.inputData.consoleInput.commandInputLines {
				inputChannel <- testData.inputData.consoleInput.commandInputLines[i]
				data := <-outputChannel
				resultData = append(resultData, data)
			}
			cancel()
			wg.Wait()

			// Assert
			assert.Equal(t, testData.expectedResult.outputChannelData, resultData)
		})
	}
}

// Test_ProcessorError creates processor, starts command processing and checks that
//
func Test_ProcessorError(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      inputData
		expectedResult expectedResult
	}{
		{
			testCaseName: unknownCommandTestCaseName,
			inputData: inputData{
				consoleInput: consoleInput{
					commandString: "unknown_command arg1 arg2\n",
				},
				parsedPipe: []string{"unknown_command arg1 arg2"},
				parsedCommands: []parsedCommand{
					{"unknown_command", "arg1", "arg2"},
				},
				processorCommands: []string{"command1"},
			},
			expectedResult: expectedResult{
				commandExecutionNumber: 0,
				outputChannelData:      []string{},
				errorChannelErrors:     []error{errorTypes.ErrorCommandNotFound},
			},
		},
		{
			testCaseName: unknownCommandInPipeTestCaseName,
			inputData: inputData{
				consoleInput: consoleInput{
					commandString: "command1 | unknown_command arg1 arg2 | command2\n",
				},
				parsedPipe: []string{"unknown_command arg1 arg2"},
				parsedCommands: []parsedCommand{
					{"unknown_command", "arg1", "arg2"},
				},
				processorCommands: []string{"command1", "command2"},
			},
			expectedResult: expectedResult{
				commandExecutionNumber: 0,
				outputChannelData:      []string{},
				errorChannelErrors:     []error{errorTypes.ErrorCommandNotFound},
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			// Prepare parsers mocks
			pipeParserMock := mocks.NewParserMock()
			commandParserMock := mocks.NewParserMock()
			pipeParserMock.On("Parse", mock.AnythingOfType("string")).Return(testData.inputData.parsedPipe)
			commandParserMock.On("Parse", mock.AnythingOfType("string")).Return(func(commandString string) []string {
				for _, parsedCommand := range testData.inputData.parsedCommands {
					if strings.HasPrefix(commandString, parsedCommand[0]) {
						return parsedCommand
					}
				}
				return []string{}
			})
			// Prepare channels
			inputChannel := make(chan string)
			outputChannel := make(chan string)
			errorChannel := make(chan error)
			// Prepare commandCreators
			commandCreatorsMap := make(map[string]commandCreator)
			commandCreator := func(
				args []string,
				commandInputChannel <-chan string,
				commandOutputChannel chan<- string,
				commandErrorChannel chan<- error,
			) command {
				return &testCommand{}
			}
			for _, processorCommand := range testData.inputData.processorCommands {
				commandCreatorsMap[processorCommand] = commandCreator
			}
			// Create processor
			processor := NewProcessor(
				pipeParserMock,
				commandParserMock,
				inputChannel,
				outputChannel,
				errorChannel,
				commandCreatorsMap,
			)
			mainCtx := context.Background()
			ctx, cancel := context.WithCancel(mainCtx)
			var wg sync.WaitGroup

			// Start processing and send command string
			wg.Add(1)
			processor.StartProcessing(ctx, &wg)
			inputChannel <- testData.inputData.consoleInput.commandString

			resultErrors := make([]error, 0)
			for range testData.expectedResult.errorChannelErrors {
				err := <-errorChannel
				resultErrors = append(resultErrors, err)
			}
			cancel()
			wg.Wait()

			// Assert
			assert.Equal(t, testData.expectedResult.errorChannelErrors, resultErrors)
		})
	}
}

// Test_ProcessorError creates processor, starts command processing and checks that
//
func Test_ProcessorCommandError(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      inputData
		expectedResult expectedResult
	}{
		{
			testCaseName: commandExecutionErrorTestCaseName,
			inputData: inputData{
				consoleInput: consoleInput{
					commandString: "command1\n",
				},
				parsedPipe: []string{"command1"},
				parsedCommands: []parsedCommand{
					{"command1"},
				},
				processorCommands:     []string{"command1"},
				commandExecutionError: commandError,
			},
			expectedResult: expectedResult{
				commandExecutionNumber: 1,
				outputChannelData:      []string{},
				errorChannelErrors:     []error{commandError},
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			// Prepare parsers mocks
			pipeParserMock := mocks.NewParserMock()
			commandParserMock := mocks.NewParserMock()
			pipeParserMock.On("Parse", mock.AnythingOfType("string")).Return(testData.inputData.parsedPipe)
			commandParserMock.On("Parse", mock.AnythingOfType("string")).Return(func(commandString string) []string {
				for _, parsedCommand := range testData.inputData.parsedCommands {
					if strings.HasPrefix(commandString, parsedCommand[0]) {
						return parsedCommand
					}
				}
				return []string{}
			})
			// Prepare channels
			inputChannel := make(chan string)
			outputChannel := make(chan string)
			errorChannel := make(chan error)
			// Prepare commandCreators
			commandCreatorsMap := make(map[string]commandCreator)
			commandCreator := func(
				args []string,
				commandInputChannel <-chan string,
				commandOutputChannel chan<- string,
				commandErrorChannel chan<- error,
			) command {
				return &testCommand{
					args:            args,
					inputChannel:    commandInputChannel,
					outputChannel:   commandOutputChannel,
					errorChannel:    commandErrorChannel,
					executionNumber: 0,
					executionTestFunc: func(
						commandInputChannel <-chan string,
						commandOutputChannel chan<- string,
						commandErrorChannel chan<- error,
						executionNumber int,
					) {
						executionNumber++
						assert.Equal(t, testData.expectedResult.commandExecutionNumber, executionNumber)
						commandErrorChannel <- testData.inputData.commandExecutionError
					},
				}
			}
			for _, processorCommand := range testData.inputData.processorCommands {
				commandCreatorsMap[processorCommand] = commandCreator
			}
			// Create processor
			processor := NewProcessor(
				pipeParserMock,
				commandParserMock,
				inputChannel,
				outputChannel,
				errorChannel,
				commandCreatorsMap,
			)
			mainCtx := context.Background()
			ctx, cancel := context.WithCancel(mainCtx)
			var wg sync.WaitGroup

			// Start processing and send command string
			wg.Add(1)
			processor.StartProcessing(ctx, &wg)
			inputChannel <- testData.inputData.consoleInput.commandString

			resultErrors := make([]error, 0)
			for range testData.expectedResult.errorChannelErrors {
				err := <-errorChannel
				resultErrors = append(resultErrors, err)
			}
			cancel()
			wg.Wait()

			// Assert
			assert.Equal(t, testData.expectedResult.errorChannelErrors, resultErrors)
		})
	}
}
