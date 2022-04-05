package command

import (
	"context"
	"errors"
	errorTypes "github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	execCommandTestCaseName                        = "Exec"
	execCommandExecErrorTestCaseName               = "Error in exec"
	execCommandNotEnoughArgumentsErrorTestCaseName = "Not enough arguments"
)

var execError = errors.New("Error in exec")

type execCommandInputData struct {
	args      []string
	execError error
}

type execCommandExpectedResult struct {
	setArgsError      error
	executableInExec  string
	argsInExec        []string
	execNumberOfCalls int
	errorChannelError error
}

func Test_ExecCommand(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      execCommandInputData
		expectedResult execCommandExpectedResult
	}{
		{
			testCaseName: execCommandTestCaseName,
			inputData: execCommandInputData{
				args:      []string{"executable", "arg1", "arg2"},
				execError: nil,
			},
			expectedResult: execCommandExpectedResult{
				executableInExec:  "executable",
				argsInExec:        []string{"arg1", "arg2"},
				execNumberOfCalls: 1,
				errorChannelError: nil,
			},
		},
		{
			testCaseName: execCommandExecErrorTestCaseName,
			inputData: execCommandInputData{
				args:      []string{"executable", "arg1", "arg2"},
				execError: execError,
			},
			expectedResult: execCommandExpectedResult{
				executableInExec:  "executable",
				argsInExec:        []string{"arg1", "arg2"},
				execNumberOfCalls: 1,
				errorChannelError: execError,
			},
		},
		{
			testCaseName: execCommandNotEnoughArgumentsErrorTestCaseName,
			inputData: execCommandInputData{
				args:      []string{},
				execError: nil,
			},
			expectedResult: execCommandExpectedResult{
				setArgsError:      errorTypes.ErrorNotEnoughArguments,
				execNumberOfCalls: 0,
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			inputChannel := make(chan string)
			outputChannel := make(chan string)
			errorChannel := make(chan error)
			execCommand := NewExecCommand(
				inputChannel,
				outputChannel,
				errorChannel,
			)
			err := execCommand.SetArgs(testData.inputData.args)
			if testData.expectedResult.setArgsError != nil {
				assert.ErrorIs(t, err, testData.expectedResult.setArgsError)
			} else {
				originalExec := exec
				execNumberOfCalls := 0
				exec = func(executable string, args []string) error {
					execNumberOfCalls++
					assert.Equal(t, testData.expectedResult.executableInExec, executable)
					assert.Equal(t, testData.expectedResult.argsInExec, args)
					return testData.inputData.execError
				}

				mainCtx := context.Background()
				ctx, _ := context.WithCancel(mainCtx)
				var wg sync.WaitGroup
				wg.Add(1)
				go execCommand.Execute(ctx, &wg)
				var resultError error
				if testData.expectedResult.errorChannelError != nil {
					resultError = <-errorChannel
				}
				wg.Wait()
				exec = originalExec

				assert.Equal(t, testData.expectedResult.errorChannelError, resultError)
				assert.Equal(t, testData.expectedResult.execNumberOfCalls, execNumberOfCalls)
			}
		})
	}
}
