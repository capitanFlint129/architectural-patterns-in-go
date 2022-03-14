package handler

import (
	"github.com/sirupsen/logrus"
)

type supportOperator struct {
	nextHandler         Handler
	problemsToSolutions map[string]string
}

// Handle tries process request or transfers it to next handler
func (s *supportOperator) Handle(problem string, logger *logrus.Logger) (string, error) {
	if solution, ok := s.problemsToSolutions[problem]; ok {
		logger.Info("Operator processed request")
		return solution, nil
	} else {
		logger.Info("Operator transferred request to engineer")
		solution, err := s.nextHandler.Handle(problem, logger)
		return solution, err
	}
}

// NewSupportOperator creates new support operator
func NewSupportOperator(nextHandler Handler, problemsToSolutions map[string]string) Handler {
	return &supportOperator{
		nextHandler:         nextHandler,
		problemsToSolutions: problemsToSolutions,
	}
}
