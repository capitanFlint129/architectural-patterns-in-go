package handler

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/chain_of_responsibility/handler/mocks"
)

type inputDataSupportRobot struct {
	request             string
	problemsToSolutions map[string]string
	handlerCreator      func(nextHandler Handler, problemsToSolutions map[string]string, logger *logrus.Logger) Handler
}

type expectedResultSupportRobot struct {
	solution               string
	error                  error
	timesNextHandlerCalled int
}

var (
	robotCanProcessRequestTestCaseName  = "Support robot can process request"
	robotCantProcessRequestTestCaseName = "Support robot can't process request"

	loggerSupportRobot = logrus.New()
)

func Test_SupportRobot(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      inputDataSupportRobot
		expectedResult expectedResultSupportRobot
	}{
		{
			testCaseName: robotCanProcessRequestTestCaseName,
			inputData: inputDataSupportRobot{
				request: "laptop broken",
				problemsToSolutions: map[string]string{
					"laptop broken": "dance with a tambourine",
				},
				handlerCreator: NewSupportRobot,
			},
			expectedResult: expectedResultSupportRobot{
				solution:               "dance with a tambourine",
				error:                  nil,
				timesNextHandlerCalled: 0,
			},
		},
		{
			testCaseName: robotCantProcessRequestTestCaseName,
			inputData: inputDataSupportRobot{
				request:             "laptop broken",
				problemsToSolutions: map[string]string{},
				handlerCreator:      NewSupportRobot,
			},
			expectedResult: expectedResultSupportRobot{
				solution:               "",
				error:                  nil,
				timesNextHandlerCalled: 1,
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			nextHandler := mocks.NewHandler()
			nextHandler.On("Handle", testData.inputData.request).Return(testData.expectedResult.solution, testData.expectedResult.error)
			handler := testData.inputData.handlerCreator(nextHandler, testData.inputData.problemsToSolutions, loggerSupportRobot)

			solution, err := handler.Handle(testData.inputData.request)
			assert.ErrorIs(t, err, testData.expectedResult.error)
			assert.Equal(t, solution, testData.expectedResult.solution)

			nextHandler.AssertNumberOfCalls(t, "Handle", testData.expectedResult.timesNextHandlerCalled)
		})
	}
}
