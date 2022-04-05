package command

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	errorTypes "github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
)

const (
	psCommandTestCaseName            = "Get processes"
	psCommandErrorInPsTestCaseName   = "Error in ps"
	psCommandTooManyArgsTestCaseName = "Too many args"
)

var psError = errors.New("error in ps")

type psCommandInputData struct {
	args      []string
	processes []process
	psError   error
}

type psCommandExpectedResult struct {
	setArgsError      error
	outputChannelData []string
	psNumberOfCalls   int
	errorChannelError error
}

func Test_PsCommand(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      psCommandInputData
		expectedResult psCommandExpectedResult
	}{
		{
			testCaseName: psCommandTestCaseName,
			inputData: psCommandInputData{
				processes: []process{
					{
						pid:        0,
						executable: "executable0",
					},
					{
						pid:        1,
						executable: "executable1",
					},
				},
			},
			expectedResult: psCommandExpectedResult{
				outputChannelData: []string{
					"0 - executable0",
					"1 - executable1",
				},
				psNumberOfCalls:   1,
				errorChannelError: nil,
			},
		},
		{
			testCaseName: psCommandErrorInPsTestCaseName,
			inputData: psCommandInputData{
				psError: psError,
			},
			expectedResult: psCommandExpectedResult{
				outputChannelData: []string{},
				psNumberOfCalls:   1,
				errorChannelError: psError,
			},
		},
		{
			testCaseName: psCommandTooManyArgsTestCaseName,
			inputData: psCommandInputData{
				args: []string{"arg1"},
			},
			expectedResult: psCommandExpectedResult{
				outputChannelData: []string{},
				psNumberOfCalls:   0,
				setArgsError:      errorTypes.ErrorTooManyArguments,
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			inputChannel := make(chan string)
			outputChannel := make(chan string)
			errorChannel := make(chan error)
			psCommand := NewPsCommand(
				inputChannel,
				outputChannel,
				errorChannel,
			)
			err := psCommand.SetArgs(testData.inputData.args)
			if testData.expectedResult.setArgsError != nil {
				assert.ErrorIs(t, err, testData.expectedResult.setArgsError)
			} else {
				originalPs := ps
				psCallsNumber := 0
				ps = func() ([]process, error) {
					psCallsNumber++
					return testData.inputData.processes, testData.inputData.psError
				}

				mainCtx := context.Background()
				ctx, _ := context.WithCancel(mainCtx)
				var wg sync.WaitGroup
				wg.Add(1)
				go psCommand.Execute(ctx, &wg)
				var resultError error
				if testData.expectedResult.errorChannelError != nil {
					resultError = <-errorChannel
				}
				outputData := make([]string, len(testData.expectedResult.outputChannelData))
				for i := range testData.expectedResult.outputChannelData {
					outputData[i] = <-outputChannel
				}
				wg.Wait()
				ps = originalPs

				assert.Equal(t, testData.expectedResult.outputChannelData, outputData)
				assert.Equal(t, testData.expectedResult.errorChannelError, resultError)
				assert.Equal(t, testData.expectedResult.psNumberOfCalls, psCallsNumber)
			}
		})
	}
}
