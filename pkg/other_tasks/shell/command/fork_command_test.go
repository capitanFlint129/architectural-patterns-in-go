package command

import (
	"errors"
	errorTypes "github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	forkCommandTestCaseName                        = "Fork"
	forkCommandForkErrorTestCaseName               = "Error in fork"
	forkCommandNotEnoughArgumentsErrorTestCaseName = "Not enough arguments"
)

var forkError = errors.New("Error in fork")

type forkCommandInputData struct {
	args        []string
	returnedPid int
	forkError   error
}

type forkCommandExpectedResult struct {
	executableInFork  string
	argsInFork        []string
	forkNumberOfCalls int
	errorChannelError error
}

func Test_ForkCommand(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      forkCommandInputData
		expectedResult forkCommandExpectedResult
	}{
		{
			testCaseName: forkCommandTestCaseName,
			inputData: forkCommandInputData{
				args:      []string{"executable", "arg1", "arg2"},
				forkError: nil,
			},
			expectedResult: forkCommandExpectedResult{
				executableInFork:  "executable",
				argsInFork:        []string{"arg0", "arg1"},
				forkNumberOfCalls: 1,
				errorChannelError: nil,
			},
		},
		{
			testCaseName: forkCommandForkErrorTestCaseName,
			inputData: forkCommandInputData{
				args:      []string{"executable", "arg1", "arg2"},
				forkError: forkError,
			},
			expectedResult: forkCommandExpectedResult{
				executableInFork:  "executable",
				argsInFork:        []string{"arg0", "arg1"},
				forkNumberOfCalls: 1,
				errorChannelError: forkError,
			},
		},
		{
			testCaseName: forkCommandNotEnoughArgumentsErrorTestCaseName,
			inputData: forkCommandInputData{
				args:      []string{},
				forkError: forkError,
			},
			expectedResult: forkCommandExpectedResult{
				forkNumberOfCalls: 0,
				errorChannelError: errorTypes.ErrorNotEnoughArguments,
			},
		},
		{
			testCaseName: forkCommandNotEnoughArgumentsErrorTestCaseName,
			inputData: forkCommandInputData{
				args:      []string{},
				forkError: nil,
			},
			expectedResult: forkCommandExpectedResult{
				forkNumberOfCalls: 0,
				errorChannelError: errorTypes.ErrorTooManyArguments,
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			inputChannel := make(chan string)
			outputChannel := make(chan string)
			errorChannel := make(chan error)

			forkCommand := NewForkCommand(
				testData.inputData.args,
				inputChannel,
				outputChannel,
				errorChannel,
			)

			originalFork := fork
			forkNumberOfCalls := 0
			fork = func(executable string, args []string) (int, error) {
				forkNumberOfCalls++
				assert.Equal(t, testData.expectedResult.executableInFork, executable)
				assert.Equal(t, testData.expectedResult.argsInFork, args)
				return testData.inputData.returnedPid, testData.inputData.forkError
			}

			var wg sync.WaitGroup
			wg.Add(1)
			go forkCommand.Execute(&wg)
			var resultError error
			if testData.expectedResult.errorChannelError != nil {
				resultError = <-errorChannel
			}
			wg.Wait()
			fork = originalFork

			assert.Equal(t, testData.expectedResult.errorChannelError, resultError)
			assert.Equal(t, testData.expectedResult.forkNumberOfCalls, forkNumberOfCalls)
		})
	}
}
