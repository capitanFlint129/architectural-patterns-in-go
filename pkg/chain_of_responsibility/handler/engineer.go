package handler

import (
	"github.com/sirupsen/logrus"

	errTypes "github.com/capitanFlint129/architectural-patterns-in-go/pkg/chain_of_responsibility/error"
)

type engineer struct {
	problemsToSolutions map[string]string
}

// Handle tries process request or transfers it to next handler
func (e *engineer) Handle(problem string, logger *logrus.Logger) (string, error) {
	if solution, ok := e.problemsToSolutions[problem]; ok {
		logger.Info("Engineer processed request")
		return solution, nil
	}
	return "", errTypes.ErrorSolutionNotFound
}

// NewEngineer creates new engineer
func NewEngineer(problemsToSolutions map[string]string) Handler {
	return &engineer{
		problemsToSolutions: problemsToSolutions,
	}
}
