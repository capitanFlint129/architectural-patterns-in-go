package handler

import (
	"github.com/sirupsen/logrus"
)

type supportRobot struct {
	nextHandler         Handler
	problemsToSolutions map[string]string
}

// Handle tries process request or transfers it to next handler
func (s *supportRobot) Handle(problem string, logger *logrus.Logger) (string, error) {
	if solution, ok := s.problemsToSolutions[problem]; ok {
		logger.Info("Robot processed request")
		return solution, nil
	} else {
		logger.Info("Robot transferred request to support operator")
		solution, err := s.nextHandler.Handle(problem, logger)
		return solution, err
	}
}

// NewSupportRobot creates new support robot
func NewSupportRobot(nextHandler Handler, problemsToSolutions map[string]string) Handler {
	return &supportRobot{
		nextHandler:         nextHandler,
		problemsToSolutions: problemsToSolutions,
	}
}
