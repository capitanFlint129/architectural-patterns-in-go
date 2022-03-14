package handler

import (
	"github.com/sirupsen/logrus"

	errTypes "github.com/capitanFlint129/architectural-patterns-in-go/pkg/chain_of_responsibility/error"
)

type engineer struct {
	problemsToSolutions map[string]string
	logger              *logrus.Logger
}

// Handle tries process request or transfers it to next handler
func (e *engineer) Handle(problem string) (string, error) {
	if solution, ok := e.problemsToSolutions[problem]; ok {
		e.logger.Info("Engineer processed request")
		return solution, nil
	}
	return "", errTypes.ErrorSolutionNotFound
}

// NewEngineer creates new engineer
func NewEngineer(problemsToSolutions map[string]string, logger *logrus.Logger) Handler {
	return &engineer{
		problemsToSolutions: problemsToSolutions,
		logger:              logger,
	}
}
