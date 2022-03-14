package handler

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/chain_of_responsibility/handler/mocks"
)

type inputDataSupportOperator struct {
	request             string
	problemsToSolutions map[string]string
	handlerCreator      func(nextHandler Handler, problemsToSolutions map[string]string, logger *logrus.Logger) Handler
}

type expectedResultSupportOperator struct {
	solution               string
	error                  error
	timesNextHandlerCalled int
}

var (
	operatorCanProcessRequestTestCaseName  = "Support operator can process request"
	operatorCantProcessRequestTestCaseName = "Support operator can't process request"

	loggerSupportOperator = logrus.New()
)

func Test_SupportOperator(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      inputDataSupportOperator
		expectedResult expectedResultSupportOperator
	}{
		{
			testCaseName: operatorCanProcessRequestTestCaseName,
			inputData: inputDataSupportOperator{
				request: "laptop broken",
				problemsToSolutions: map[string]string{
					"laptop broken": "dance with a tambourine",
				},
				handlerCreator: NewSupportOperator,
			},
			expectedResult: expectedResultSupportOperator{
				solution:               "dance with a tambourine",
				error:                  nil,
				timesNextHandlerCalled: 0,
			},
		},
		{
			testCaseName: operatorCantProcessRequestTestCaseName,
			inputData: inputDataSupportOperator{
				request:             "laptop broken",
				problemsToSolutions: map[string]string{},
				handlerCreator:      NewSupportOperator,
			},
			expectedResult: expectedResultSupportOperator{
				solution:               "",
				error:                  nil,
				timesNextHandlerCalled: 1,
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			nextHandler := mocks.NewHandler()
			nextHandler.On("Handle", testData.inputData.request).Return(testData.expectedResult.solution, testData.expectedResult.error)
			handler := testData.inputData.handlerCreator(nextHandler, testData.inputData.problemsToSolutions, loggerSupportOperator)

			solution, err := handler.Handle(testData.inputData.request)
			assert.ErrorIs(t, err, testData.expectedResult.error)
			assert.Equal(t, solution, testData.expectedResult.solution)

			nextHandler.AssertNumberOfCalls(t, "Handle", testData.expectedResult.timesNextHandlerCalled)
		})
	}
}
