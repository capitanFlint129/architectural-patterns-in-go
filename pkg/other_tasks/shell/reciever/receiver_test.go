package reciever

import (
	"bufio"
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

type inputData struct {
	data string
}

type expectedResult struct {
	outputChannelDataLines []string
}

const (
	receiveDataTestCaseName = "Receive data"
)

func Test_Receiver(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      inputData
		expectedResult expectedResult
	}{
		{
			testCaseName: receiveDataTestCaseName,
			inputData: inputData{
				data: "line1\nline2\nline3\n",
			},
			expectedResult: expectedResult{
				outputChannelDataLines: []string{"line1", "line2", "line3"},
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			var stdin bytes.Buffer
			stdin.Write([]byte(testData.inputData.data))
			scanner := bufio.NewScanner(&stdin)
			outputChannel := make(chan string)
			errorChannel := make(chan error)
			receiver := NewReceiver(scanner, outputChannel, errorChannel)
			mainCtx := context.Background()
			ctx, cancel := context.WithCancel(mainCtx)
			var wg sync.WaitGroup

			wg.Add(1)
			receiver.StartReceive(ctx, &wg)
			outputChannelLines := make([]string, len(testData.expectedResult.outputChannelDataLines))
			for i := 0; i < len(testData.expectedResult.outputChannelDataLines); i++ {
				outputChannelLines[i] = <-outputChannel
			}
			cancel()

			wg.Wait() // TODO вопрос: wg в тесте?
			assert.Equal(t, testData.expectedResult.outputChannelDataLines, outputChannelLines)
		})
	}
}
