package responder

import (
	"bytes"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

const (
	respondTestCaseName = "Respond"
)

type inputData struct {
	inputChannelData []string
	errorChannelData []string
}

type expectedResult struct {
	outputData string
	errorData  string
}

func Test_Responder(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      inputData
		expectedResult expectedResult
	}{
		{
			testCaseName: respondTestCaseName,
			inputData: inputData{
				inputChannelData: []string{
					"line1",
					"line2",
					"line3",
				},
				errorChannelData: []string{
					"err1",
					"err2",
					"err3",
				},
			},
			expectedResult: expectedResult{
				outputData: "line1\nline2\nline3\n",
				errorData:  "err1\nerr2\nerr3\n",
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			inputChannel := make(chan string)
			errorChannel := make(chan error)
			var outputWriter bytes.Buffer
			var errorWriter bytes.Buffer
			responder := NewResponder(&outputWriter, &errorWriter, inputChannel, errorChannel)
			mainCtx := context.Background()
			ctx, cancel := context.WithCancel(mainCtx)
			var wg sync.WaitGroup

			wg.Add(1)
			responder.StartRespond(ctx, &wg)
			for _, data := range testData.inputData.inputChannelData {
				inputChannel <- data
			}
			for _, data := range testData.inputData.errorChannelData {
				err := errors.New(data)
				errorChannel <- err
			}
			cancel()

			wg.Wait()
			assert.Equal(t, testData.expectedResult.outputData, outputWriter.String())
			assert.Equal(t, testData.expectedResult.errorData, errorWriter.String())
		})
	}
}
