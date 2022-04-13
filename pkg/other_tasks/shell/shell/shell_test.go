package shell

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/shell/mocks"
	"github.com/stretchr/testify/mock"
	"sync"
	"testing"
)

const (
	shellTestCaseName = "Run shell"
)

func Test_Receiver(t *testing.T) {
	for _, testData := range []struct {
		testCaseName string
	}{
		{
			testCaseName: shellTestCaseName,
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			receiver := mocks.NewReceiver()
			processor := mocks.NewProcessor()
			responder := mocks.NewResponder()

			receiver.On("StartReceive", mock.AnythingOfType("context.Context"), mock.AnythingOfType("*sync.WaitGroup")).Run(func(args mock.Arguments) {
				args.Get(1).(*sync.WaitGroup).Done()
			})
			processor.On("StartProcessing", mock.AnythingOfType("context.Context"), mock.AnythingOfType("sync.WaitGroup")).Run(func(args mock.Arguments) {
				args.Get(1).(*sync.WaitGroup).Done()
			})
			responder.On("StartRespond", mock.AnythingOfType("context.Context"), mock.AnythingOfType("sync.WaitGroup")).Run(func(args mock.Arguments) {
				args.Get(1).(*sync.WaitGroup).Done()
			})

			shell := NewShell(receiver, processor, responder)
			mainCtx := context.Background()
			ctx, _ := context.WithCancel(mainCtx)
			shell.Run(ctx)
		})
	}
}
