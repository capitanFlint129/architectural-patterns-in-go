package main

import (
	"github.com/sirupsen/logrus"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/chain_of_responsibility/handler"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/chain_of_responsibility/support"
)

var (
	robotProblemsToSolution = map[string]string{
		"laptop freezed": "restart it",
	}
	operatorProblemsToSolution = map[string]string{
		"laptop doesn't turn on": "bring it to a service center",
	}
	engineerProblemsToSolution = map[string]string{
		"laptop does not see printer": "download driver",
	}
)

func main() {
	logger := logrus.New()

	// Подготовка цепочки обязанностей
	engineer := handler.NewEngineer(engineerProblemsToSolution, logger)
	operator := handler.NewSupportOperator(engineer, operatorProblemsToSolution, logger)
	robot := handler.NewSupportRobot(operator, robotProblemsToSolution, logger)

	dellSupport := support.NewSupport([]handler.Handler{robot, operator, engineer}, logger)
	solution, err := dellSupport.ProcessRequest("laptop doesn't turn on")
	if err != nil {
		logger.Errorf("Can't process request: %s", err.Error())
	}
	logger.Info(solution)
}
