package command

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	pwdCommandTestCaseName           = "Print working directory"
	pwdCommandGetwdErrorTestCaseName = "Error in getwd"
)

var getwdError = errors.New("Error in getwd")

type pwdCommandInputData struct {
	args               []string
	getwdError         error
	returnedWorkingDir string
}

type pwdCommandExpectedResult struct {
	getwdNumberOfCalls int
	outputChannelData  string
	errorChannelError  error
}

func Test_PwdCommand(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      pwdCommandInputData
		expectedResult pwdCommandExpectedResult
	}{
		{
			testCaseName: pwdCommandTestCaseName,
			inputData: pwdCommandInputData{
				args:               []string{},
				returnedWorkingDir: "working_directory",
				getwdError:         nil,
			},
			expectedResult: pwdCommandExpectedResult{
				getwdNumberOfCalls: 1,
				outputChannelData:  "working_directory",
				errorChannelError:  nil,
			},
		},
		{
			testCaseName: pwdCommandGetwdErrorTestCaseName,
			inputData: pwdCommandInputData{
				args:               []string{},
				returnedWorkingDir: "working_directory",
				getwdError:         getwdError,
			},
			expectedResult: pwdCommandExpectedResult{
				getwdNumberOfCalls: 1,
				outputChannelData:  "",
				errorChannelError:  getwdError,
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			inputChannel := make(chan string)
			outputChannel := make(chan string)
			errorChannel := make(chan error)
			pwdCommand := NewPwdCommand(
				inputChannel,
				outputChannel,
				errorChannel,
			)
			_ = pwdCommand.SetArgs(testData.inputData.args)

			originalGetwd := getwd
			getwdNumberOfCalls := 0
			getwd = func() (string, error) {
				getwdNumberOfCalls++
				return testData.inputData.returnedWorkingDir, testData.inputData.getwdError
			}

			mainCtx := context.Background()
			ctx, _ := context.WithCancel(mainCtx)
			var wg sync.WaitGroup
			wg.Add(1)
			go pwdCommand.Execute(ctx, &wg)
			var resultData string
			var resultError error
			select {
			case resultData = <-outputChannel:
			case resultError = <-errorChannel:
			}
			wg.Wait()
			getwd = originalGetwd

			assert.Equal(t, testData.expectedResult.outputChannelData, resultData)
			assert.Equal(t, testData.expectedResult.errorChannelError, resultError)
			assert.Equal(t, testData.expectedResult.getwdNumberOfCalls, getwdNumberOfCalls)
		})
	}
}
