package handler

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	errTypes "github.com/capitanFlint129/architectural-patterns-in-go/pkg/chain_of_responsibility/error"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/chain_of_responsibility/support/mocks"
)

type inputData struct {
	request             string
	problemsToSolutions map[string]string
	handlerCreator      func(nextHandler Handler, problemsToSolutions map[string]string) Handler
}

type expectedResult struct {
	solution               string
	error                  error
	timesNextHandlerCalled int
}

func Test_Handler(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      inputData
		expectedResult expectedResult
	}{
		{
			testCaseName: "Support robot can process request",
			inputData: inputData{
				request: "laptop broken",
				problemsToSolutions: map[string]string{
					"laptop broken": "dance with a tambourine",
				},
				handlerCreator: NewSupportRobot,
			},
			expectedResult: expectedResult{
				solution:               "dance with a tambourine",
				error:                  nil,
				timesNextHandlerCalled: 0,
			},
		},
		{
			testCaseName: "Support robot can't process request",
			inputData: inputData{
				request:             "laptop broken",
				problemsToSolutions: map[string]string{},
				handlerCreator:      NewSupportRobot,
			},
			expectedResult: expectedResult{
				solution:               "",
				error:                  nil,
				timesNextHandlerCalled: 1,
			},
		},
		{
			testCaseName: "Support operator can process request",
			inputData: inputData{
				request: "laptop broken",
				problemsToSolutions: map[string]string{
					"laptop broken": "dance with a tambourine",
				},
				handlerCreator: NewSupportOperator,
			},
			expectedResult: expectedResult{
				solution:               "dance with a tambourine",
				error:                  nil,
				timesNextHandlerCalled: 0,
			},
		},
		{
			testCaseName: "Support operator can't process request",
			inputData: inputData{
				request:             "laptop broken",
				problemsToSolutions: map[string]string{},
				handlerCreator:      NewSupportOperator,
			},
			expectedResult: expectedResult{
				solution:               "",
				error:                  nil,
				timesNextHandlerCalled: 1,
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			logger := logrus.New()
			nextHandler := mocks.NewHandler()
			nextHandler.On("Handle", testData.inputData.request, logger).Return(testData.expectedResult.solution, testData.expectedResult.error)

			handler := testData.inputData.handlerCreator(nextHandler, testData.inputData.problemsToSolutions)
			solution, err := handler.Handle(testData.inputData.request, logger)

			assert.Equal(t, solution, testData.expectedResult.solution)
			assert.ErrorIs(t, err, testData.expectedResult.error)
			nextHandler.AssertNumberOfCalls(t, "Handle", testData.expectedResult.timesNextHandlerCalled)
		})
	}
}

type inputDataFinalHandler struct {
	request             string
	problemsToSolutions map[string]string
	handlerCreator      func(problemsToSolutions map[string]string) Handler
}

func Test_FinalHandler(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      inputDataFinalHandler
		expectedResult expectedResult
	}{
		{
			testCaseName: "Engineer can process request",
			inputData: inputDataFinalHandler{
				handlerCreator: NewEngineer,
				request:        "laptop broken",
				problemsToSolutions: map[string]string{
					"laptop broken": "dance with a tambourine",
				},
			},
			expectedResult: expectedResult{
				solution: "dance with a tambourine",
				error:    nil,
			},
		},
		{
			testCaseName: "Engineer can't process request",
			inputData: inputDataFinalHandler{
				handlerCreator:      NewEngineer,
				request:             "laptop broken",
				problemsToSolutions: map[string]string{},
			},
			expectedResult: expectedResult{
				solution: "",
				error:    errTypes.ErrorSolutionNotFound,
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			logger := logrus.New()

			handler := testData.inputData.handlerCreator(testData.inputData.problemsToSolutions)
			solution, err := handler.Handle(testData.inputData.request, logger)

			assert.Equal(t, solution, testData.expectedResult.solution)
			assert.ErrorIs(t, err, testData.expectedResult.error)
		})
	}
}
