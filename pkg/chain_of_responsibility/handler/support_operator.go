package handler

import (
	"github.com/sirupsen/logrus"
)

type supportOperator struct {
	nextHandler         Handler
	problemsToSolutions map[string]string
	logger              *logrus.Logger
}

// Handle tries process request or transfers it to next handler
func (s *supportOperator) Handle(problem string) (string, error) {
	if solution, ok := s.problemsToSolutions[problem]; ok {
		s.logger.Info("Operator processed request")
		return solution, nil
	} else {
		s.logger.Info("Operator transferred request to engineer")
		solution, err := s.nextHandler.Handle(problem)
		return solution, err
	}
}

// NewSupportOperator creates new support operator
func NewSupportOperator(nextHandler Handler, problemsToSolutions map[string]string, logger *logrus.Logger) Handler {
	return &supportOperator{
		nextHandler:         nextHandler,
		problemsToSolutions: problemsToSolutions,
		logger:              logger,
	}
}
