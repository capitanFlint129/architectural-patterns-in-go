package handler

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	errTypes "github.com/capitanFlint129/architectural-patterns-in-go/pkg/chain_of_responsibility/error"
)

type inputDataEngineer struct {
	request             string
	problemsToSolutions map[string]string
	handlerCreator      func(problemsToSolutions map[string]string, logger *logrus.Logger) Handler
}

type expectedResultEngineer struct {
	solution               string
	error                  error
	timesNextHandlerCalled int
}

var (
	engineerCanProcessRequestTestCaseName  = "Engineer can process request"
	engineerCantProcessRequestTestCaseName = "Engineer can't process request"

	loggerEngineer = logrus.New()
)

func Test_FinalHandler(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      inputDataEngineer
		expectedResult expectedResultEngineer
	}{
		{
			testCaseName: engineerCanProcessRequestTestCaseName,
			inputData: inputDataEngineer{
				handlerCreator: NewEngineer,
				request:        "laptop broken",
				problemsToSolutions: map[string]string{
					"laptop broken": "dance with a tambourine",
				},
			},
			expectedResult: expectedResultEngineer{
				solution: "dance with a tambourine",
				error:    nil,
			},
		},
		{
			testCaseName: engineerCantProcessRequestTestCaseName,
			inputData: inputDataEngineer{
				handlerCreator:      NewEngineer,
				request:             "laptop broken",
				problemsToSolutions: map[string]string{},
			},
			expectedResult: expectedResultEngineer{
				solution: "",
				error:    errTypes.ErrorSolutionNotFound,
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			handler := testData.inputData.handlerCreator(testData.inputData.problemsToSolutions, loggerEngineer)

			solution, err := handler.Handle(testData.inputData.request)
			assert.ErrorIs(t, err, testData.expectedResult.error)
			assert.Equal(t, solution, testData.expectedResult.solution)
		})
	}
}
