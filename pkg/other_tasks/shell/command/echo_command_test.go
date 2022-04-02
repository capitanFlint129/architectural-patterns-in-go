package command

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	echoArgsTestCaseName   = "Echo arguments"
	echoNoArgsTestCaseName = "Echo no arguments"
)

type echoCommandInputData struct {
	args []string
}

type echoCommandExpectedResult struct {
	outputChannelData string
}

func Test_EchoCommand(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      echoCommandInputData
		expectedResult echoCommandExpectedResult
	}{
		{
			testCaseName: echoArgsTestCaseName,
			inputData: echoCommandInputData{
				args: []string{"arg1", "arg2"},
			},
			expectedResult: echoCommandExpectedResult{
				outputChannelData: "arg1 arg2",
			},
		},
		{
			testCaseName: echoNoArgsTestCaseName,
			inputData: echoCommandInputData{
				args: []string{},
			},
			expectedResult: echoCommandExpectedResult{
				outputChannelData: "",
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			inputChannel := make(chan string)
			outputChannel := make(chan string)
			errorChannel := make(chan error)

			echoCommand := NewEchoCommand(
				testData.inputData.args,
				inputChannel,
				outputChannel,
				errorChannel,
			)

			var wg sync.WaitGroup
			wg.Add(1)
			go echoCommand.Execute(&wg)
			result := <-outputChannel
			wg.Wait()

			assert.Equal(t, testData.expectedResult.outputChannelData, result)
		})
	}
}
