package handler

import (
	"github.com/sirupsen/logrus"
)

type supportRobot struct {
	nextHandler         Handler
	problemsToSolutions map[string]string
	logger              *logrus.Logger
}

// Handle tries process request or transfers it to next handler
func (s *supportRobot) Handle(problem string) (string, error) {
	if solution, ok := s.problemsToSolutions[problem]; ok {
		s.logger.Info("Robot processed request")
		return solution, nil
	} else {
		s.logger.Info("Robot transferred request to support operator")
		solution, err := s.nextHandler.Handle(problem)
		return solution, err
	}
}

// NewSupportRobot creates new support robot
func NewSupportRobot(nextHandler Handler, problemsToSolutions map[string]string, logger *logrus.Logger) Handler {
	return &supportRobot{
		nextHandler:         nextHandler,
		problemsToSolutions: problemsToSolutions,
		logger:              logger,
	}
}
