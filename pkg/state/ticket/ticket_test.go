package ticket

import (
	"testing"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/state/ticket/mocks"
	"github.com/sirupsen/logrus"
)

const (
	ticketUsageTestCaseName = "Ticket usage"
)

type expectedResult struct {
	numberOfCallsOfStateMethods int
}

var logger = logrus.New()

func Test_Ticket(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		expectedResult expectedResult
	}{
		{
			testCaseName: ticketUsageTestCaseName,
			expectedResult: expectedResult{
				numberOfCallsOfStateMethods: 1,
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			stateMock := mocks.NewStateMock()
			stateMock.On("Publish").Return()
			stateMock.On("Complete").Return()
			stateMock.On("Delete").Return()

			ticket := NewTicket(logger)
			ticket.SetState(stateMock)

			ticket.Publish()
			ticket.Complete()
			ticket.Delete()

			stateMock.AssertNumberOfCalls(t, "Publish", testData.expectedResult.numberOfCallsOfStateMethods)
			stateMock.AssertNumberOfCalls(t, "Complete", testData.expectedResult.numberOfCallsOfStateMethods)
			stateMock.AssertNumberOfCalls(t, "Delete", testData.expectedResult.numberOfCallsOfStateMethods)
		})
	}
}
