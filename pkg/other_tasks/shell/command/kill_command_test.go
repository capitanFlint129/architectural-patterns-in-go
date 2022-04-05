package command

//import (
//	"context"
//	"errors"
//	errorTypes "github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
//	"sync"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//)
//
//const (
//	killCommandTestCaseName                        = "Kill"
//	killCommandKillErrorTestCaseName               = "Error in kill"
//	killCommandNotEnoughArgumentsErrorTestCaseName = "Not enough arguments"
//	killCommandTooManyArgumentsErrorTestCaseName   = "Too many arguments"
//)
//
//var killError = errors.New("Error in kill")
//
//type killCommandInputData struct {
//	args      []string
//	killError error
//}
//
//type killCommandExpectedResult struct {
//	pidInKill         int
//	killNumberOfCalls int
//	setArgsError      error
//	errorChannelError error
//}
//
//func Test_KillCommand(t *testing.T) {
//	for _, testData := range []struct {
//		testCaseName   string
//		inputData      killCommandInputData
//		expectedResult killCommandExpectedResult
//	}{
//		{
//			testCaseName: killCommandTestCaseName,
//			inputData: killCommandInputData{
//				args:      []string{"0"},
//				killError: nil,
//			},
//			expectedResult: killCommandExpectedResult{
//				pidInKill:         0,
//				killNumberOfCalls: 1,
//				errorChannelError: nil,
//			},
//		},
//		{
//			testCaseName: killCommandKillErrorTestCaseName,
//			inputData: killCommandInputData{
//				args:      []string{"0"},
//				killError: killError,
//			},
//			expectedResult: killCommandExpectedResult{
//				pidInKill:         0,
//				killNumberOfCalls: 1,
//				errorChannelError: killError,
//			},
//		},
//		{
//			testCaseName: killCommandNotEnoughArgumentsErrorTestCaseName,
//			inputData: killCommandInputData{
//				args:      []string{},
//				killError: nil,
//			},
//			expectedResult: killCommandExpectedResult{
//				killNumberOfCalls: 0,
//				setArgsError:      errorTypes.ErrorNotEnoughArguments,
//			},
//		},
//		{
//			testCaseName: killCommandTooManyArgumentsErrorTestCaseName,
//			inputData: killCommandInputData{
//				args:      []string{"0", "100"},
//				killError: nil,
//			},
//			expectedResult: killCommandExpectedResult{
//				killNumberOfCalls: 0,
//				setArgsError:      errorTypes.ErrorTooManyArguments,
//			},
//		},
//	} {
//		t.Run(testData.testCaseName, func(t *testing.T) {
//			inputChannel := make(chan string)
//			outputChannel := make(chan string)
//			errorChannel := make(chan error)
//			killCommand := NewKillCommand(
//				inputChannel,
//				outputChannel,
//				errorChannel,
//			)
//			err := killCommand.SetArgs(testData.inputData.args)
//			if testData.expectedResult.setArgsError != nil {
//				assert.ErrorIs(t, err, testData.expectedResult.setArgsError)
//			} else {
//
//				originalKill := kill
//				killNumberOfCalls := 0
//				kill = func(pid int) error {
//					killNumberOfCalls++
//					assert.Equal(t, testData.expectedResult.pidInKill, pid)
//					return testData.inputData.killError
//				}
//
//				mainCtx := context.Background()
//				ctx, _ := context.WithCancel(mainCtx)
//				var wg sync.WaitGroup
//				wg.Add(1)
//				go killCommand.Execute(ctx, &wg)
//				var resultError error
//				if testData.expectedResult.errorChannelError != nil {
//					resultError = <-errorChannel
//				}
//				wg.Wait()
//				kill = originalKill
//
//				assert.Equal(t, testData.expectedResult.errorChannelError, resultError)
//				assert.Equal(t, testData.expectedResult.killNumberOfCalls, killNumberOfCalls)
//			}
//		})
//	}
//}
