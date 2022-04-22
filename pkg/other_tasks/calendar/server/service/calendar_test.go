package service

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	createOneEventAndGetItTestCaseName       = "Create event and get it"
	createTreeEventsAndGetItTestCaseName     = "Create tree events and get it"
	createEventsAndGetSomeOfThemTestCaseName = "Create tree events and get it, but one out of period"
	createEventAlreadyExistsTestCaseName     = "Create event that already exists"
	updateEventAndGetItTestCaseName          = "Update event test"
	updateEventNotFoundTestCaseName          = "Update event that does not exist"
	deleteEventTestCaseName                  = "Delete event test"
	deleteEventNotFoundTestCaseName          = "Delete event that does not exist"
)

type createOrUpdateEventReturnedData struct {
	event types.Event
	err   error
}

type eventsForPeriodReturnedData struct {
	events []types.Event
	err    error
}

type createEventInputData struct {
	createEventsData    []types.EventHandlerData
	eventsForPeriodData types.DateIntervalHandlerData
}

type createEventExpectedResult struct {
	createEventsResult    []createOrUpdateEventReturnedData
	eventsForPeriodResult eventsForPeriodReturnedData
}

func TestCalendar_CreateEvent(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      createEventInputData
		expectedResult createEventExpectedResult
	}{
		{
			testCaseName: createOneEventAndGetItTestCaseName,
			inputData: createEventInputData{
				createEventsData: []types.EventHandlerData{
					{
						UserId: 0,
						Event: types.Event{
							Name: "event0",
							Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
						},
					},
				},
				eventsForPeriodData: types.DateIntervalHandlerData{
					UserId:    0,
					StartDate: time.Date(2022, time.January, 12, 0, 0, 0, 0, time.UTC),
					EndDate:   time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedResult: createEventExpectedResult{
				createEventsResult: []createOrUpdateEventReturnedData{
					{
						event: types.Event{
							Name: "event0",
							Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
						},
					},
				},
				eventsForPeriodResult: eventsForPeriodReturnedData{
					events: []types.Event{
						{
							Name: "event0",
							Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
						},
					},
				},
			},
		},
		{
			testCaseName: createTreeEventsAndGetItTestCaseName,
			inputData: createEventInputData{
				createEventsData: []types.EventHandlerData{
					{
						UserId: 0,
						Event: types.Event{
							Name: "event0",
							Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
						},
					},
					{
						UserId: 0,
						Event: types.Event{
							Name: "event1",
							Date: time.Date(2022, time.January, 14, 0, 0, 0, 0, time.UTC),
						},
					},
					{
						UserId: 0,
						Event: types.Event{
							Name: "event2",
							Date: time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC),
						},
					},
				},
				eventsForPeriodData: types.DateIntervalHandlerData{
					UserId:    0,
					StartDate: time.Date(2022, time.January, 12, 0, 0, 0, 0, time.UTC),
					EndDate:   time.Date(2022, time.January, 16, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedResult: createEventExpectedResult{
				createEventsResult: []createOrUpdateEventReturnedData{
					{
						event: types.Event{
							Name: "event0",
							Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
						},
					},
					{
						event: types.Event{
							Name: "event1",
							Date: time.Date(2022, time.January, 14, 0, 0, 0, 0, time.UTC),
						},
					},
					{
						event: types.Event{
							Name: "event2",
							Date: time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC),
						},
					},
				},
				eventsForPeriodResult: eventsForPeriodReturnedData{
					events: []types.Event{
						{
							Name: "event0",
							Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
						},
						{
							Name: "event1",
							Date: time.Date(2022, time.January, 14, 0, 0, 0, 0, time.UTC),
						},
						{
							Name: "event2",
							Date: time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC),
						},
					},
				},
			},
		},
		{
			testCaseName: createEventsAndGetSomeOfThemTestCaseName,
			inputData: createEventInputData{
				createEventsData: []types.EventHandlerData{
					{
						UserId: 0,
						Event: types.Event{
							Name: "event0",
							Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
						},
					},
					{
						UserId: 0,
						Event: types.Event{
							Name: "event1",
							Date: time.Date(2022, time.January, 14, 0, 0, 0, 0, time.UTC),
						},
					},
					{
						UserId: 0,
						Event: types.Event{
							Name: "event2",
							Date: time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC),
						},
					},
				},
				eventsForPeriodData: types.DateIntervalHandlerData{
					UserId:    0,
					StartDate: time.Date(2022, time.January, 12, 0, 0, 0, 0, time.UTC),
					EndDate:   time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedResult: createEventExpectedResult{
				createEventsResult: []createOrUpdateEventReturnedData{
					{
						event: types.Event{
							Name: "event0",
							Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
						},
					},
					{
						event: types.Event{
							Name: "event1",
							Date: time.Date(2022, time.January, 14, 0, 0, 0, 0, time.UTC),
						},
					},
					{
						event: types.Event{
							Name: "event2",
							Date: time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC),
						},
					},
				},
				eventsForPeriodResult: eventsForPeriodReturnedData{
					events: []types.Event{
						{
							Name: "event0",
							Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
						},
						{
							Name: "event1",
							Date: time.Date(2022, time.January, 14, 0, 0, 0, 0, time.UTC),
						},
					},
				},
			},
		},
		{
			testCaseName: createEventAlreadyExistsTestCaseName,
			inputData: createEventInputData{
				createEventsData: []types.EventHandlerData{
					{
						UserId: 0,
						Event: types.Event{
							Name: "event0",
							Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
						},
					},
					{
						UserId: 0,
						Event: types.Event{
							Name: "event0",
							Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
						},
					},
				},
				eventsForPeriodData: types.DateIntervalHandlerData{
					UserId:    0,
					StartDate: time.Date(2022, time.January, 12, 0, 0, 0, 0, time.UTC),
					EndDate:   time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedResult: createEventExpectedResult{
				createEventsResult: []createOrUpdateEventReturnedData{
					{
						event: types.Event{
							Name: "event0",
							Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
						},
					},
					{
						event: types.Event{},
						err:   types.ErrorEventAlreadyExists,
					},
				},
				eventsForPeriodResult: eventsForPeriodReturnedData{
					events: []types.Event{
						{
							Name: "event0",
							Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
						},
					},
				},
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			calendar := NewCalendar()
			mainCtx := context.Background()
			ctx, cancel := context.WithCancel(mainCtx)
			defer cancel()

			for i := range testData.inputData.createEventsData {
				event, err := calendar.CreateEvent(ctx, testData.inputData.createEventsData[i])
				assert.Equal(t, testData.expectedResult.createEventsResult[i].event, event)
				assert.Equal(t, testData.expectedResult.createEventsResult[i].err, err)
			}
			events, err := calendar.EventsForPeriod(ctx, testData.inputData.eventsForPeriodData)

			assert.Equal(t, testData.expectedResult.eventsForPeriodResult.events, events)
			assert.Equal(t, testData.expectedResult.eventsForPeriodResult.err, err)
		})
	}
}

type updateEventInputData struct {
	createEventsData    []types.EventHandlerData
	updateEventData     types.UpdateEventHandlerData
	eventsForPeriodData types.DateIntervalHandlerData
}

type updateEventExpectedResult struct {
	updateEventResult     createOrUpdateEventReturnedData
	eventsForPeriodResult eventsForPeriodReturnedData
}

func TestCalendar_UpdateEvent(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      updateEventInputData
		expectedResult updateEventExpectedResult
	}{
		{
			testCaseName: updateEventAndGetItTestCaseName,
			inputData: updateEventInputData{
				createEventsData: []types.EventHandlerData{
					{
						UserId: 0,
						Event: types.Event{
							Name: "event0",
							Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
						},
					},
				},
				updateEventData: types.UpdateEventHandlerData{
					UserId: 0,
					Event: types.Event{
						Name: "event0",
						Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
					},
					NewEvent: types.Event{
						Name: "event1",
						Date: time.Date(2022, time.January, 14, 0, 0, 0, 0, time.UTC),
					},
				},
				eventsForPeriodData: types.DateIntervalHandlerData{
					UserId:    0,
					StartDate: time.Date(2022, time.January, 12, 0, 0, 0, 0, time.UTC),
					EndDate:   time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedResult: updateEventExpectedResult{
				updateEventResult: createOrUpdateEventReturnedData{
					event: types.Event{
						Name: "event1",
						Date: time.Date(2022, time.January, 14, 0, 0, 0, 0, time.UTC),
					},
				},
				eventsForPeriodResult: eventsForPeriodReturnedData{
					events: []types.Event{
						{
							Name: "event1",
							Date: time.Date(2022, time.January, 14, 0, 0, 0, 0, time.UTC),
						},
					},
				},
			},
		},
		{
			testCaseName: updateEventNotFoundTestCaseName,
			inputData: updateEventInputData{
				createEventsData: []types.EventHandlerData{},
				updateEventData: types.UpdateEventHandlerData{
					UserId: 0,
					Event: types.Event{
						Name: "event0",
						Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
					},
					NewEvent: types.Event{
						Name: "event1",
						Date: time.Date(2022, time.January, 14, 0, 0, 0, 0, time.UTC),
					},
				},
				eventsForPeriodData: types.DateIntervalHandlerData{
					UserId:    0,
					StartDate: time.Date(2022, time.January, 12, 0, 0, 0, 0, time.UTC),
					EndDate:   time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedResult: updateEventExpectedResult{
				updateEventResult: createOrUpdateEventReturnedData{
					event: types.Event{},
					err:   types.ErrorEventNotFound,
				},
				eventsForPeriodResult: eventsForPeriodReturnedData{
					events: []types.Event{},
				},
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			calendar := NewCalendar()
			mainCtx := context.Background()
			ctx, cancel := context.WithCancel(mainCtx)
			defer cancel()

			for i := range testData.inputData.createEventsData {
				calendar.CreateEvent(ctx, testData.inputData.createEventsData[i])
			}
			updatedEvent, updateEventErr := calendar.UpdateEvent(ctx, testData.inputData.updateEventData)
			events, err := calendar.EventsForPeriod(ctx, testData.inputData.eventsForPeriodData)

			assert.Equal(t, testData.expectedResult.updateEventResult.event, updatedEvent)
			assert.Equal(t, testData.expectedResult.updateEventResult.err, updateEventErr)
			assert.Equal(t, testData.expectedResult.eventsForPeriodResult.events, events)
			assert.Equal(t, testData.expectedResult.eventsForPeriodResult.err, err)
		})
	}
}

type deleteEventInputData struct {
	createEventsData    []types.EventHandlerData
	deleteEventData     types.EventHandlerData
	eventsForPeriodData types.DateIntervalHandlerData
}

type deleteEventExpectedResult struct {
	deleteEventError      error
	eventsForPeriodResult eventsForPeriodReturnedData
}

func TestCalendar_DeleteEvent(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      deleteEventInputData
		expectedResult deleteEventExpectedResult
	}{
		{
			testCaseName: deleteEventTestCaseName,
			inputData: deleteEventInputData{
				createEventsData: []types.EventHandlerData{
					{
						UserId: 0,
						Event: types.Event{
							Name: "event0",
							Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
						},
					},
				},
				deleteEventData: types.EventHandlerData{
					UserId: 0,
					Event: types.Event{
						Name: "event0",
						Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
					},
				},
				eventsForPeriodData: types.DateIntervalHandlerData{
					UserId:    0,
					StartDate: time.Date(2022, time.January, 12, 0, 0, 0, 0, time.UTC),
					EndDate:   time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedResult: deleteEventExpectedResult{
				eventsForPeriodResult: eventsForPeriodReturnedData{
					events: []types.Event{},
				},
			},
		},
		{
			testCaseName: deleteEventNotFoundTestCaseName,
			inputData: deleteEventInputData{
				createEventsData: []types.EventHandlerData{},
				deleteEventData: types.EventHandlerData{
					UserId: 0,
					Event: types.Event{
						Name: "event0",
						Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
					},
				},
				eventsForPeriodData: types.DateIntervalHandlerData{
					UserId:    0,
					StartDate: time.Date(2022, time.January, 12, 0, 0, 0, 0, time.UTC),
					EndDate:   time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedResult: deleteEventExpectedResult{
				deleteEventError: types.ErrorEventNotFound,
				eventsForPeriodResult: eventsForPeriodReturnedData{
					events: []types.Event{},
				},
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			calendar := NewCalendar()
			mainCtx := context.Background()
			ctx, cancel := context.WithCancel(mainCtx)
			defer cancel()

			for i := range testData.inputData.createEventsData {
				calendar.CreateEvent(ctx, testData.inputData.createEventsData[i])
			}
			deleteEventErr := calendar.DeleteEvent(ctx, testData.inputData.deleteEventData)
			events, err := calendar.EventsForPeriod(ctx, testData.inputData.eventsForPeriodData)

			assert.Equal(t, testData.expectedResult.deleteEventError, deleteEventErr)
			assert.Equal(t, testData.expectedResult.eventsForPeriodResult.events, events)
			assert.Equal(t, testData.expectedResult.eventsForPeriodResult.err, err)
		})
	}
}
