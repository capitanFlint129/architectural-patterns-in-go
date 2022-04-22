package tests

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/client/client"
	clientTransport "github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/client/transport"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/server/handler"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/server/service"
	serverTransport "github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/server/transport"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/tests/mocks"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
)

const (
	createEventTestCaseName                 = "Send create event request"
	createEventServiceErrorTestCaseName     = "Send create event request and service error occurred"
	updateEventTestCaseName                 = "Send update event request"
	updateEventServiceErrorTestCaseName     = "Send update event request and service error occurred"
	deleteEventTestCaseName                 = "Send delete event request"
	deleteEventServiceErrorTestCaseName     = "Send delete event request and service error occurred"
	eventsForPeriodTestCaseName             = "Get events for period"
	eventsForPeriodServiceErrorTestCaseName = "Get events for period and service error occurred"

	dateFormat = "2006-01-02"

	createEventHttpMethod     = http.MethodPost
	updateEventHttpMethod     = http.MethodPut
	deleteEventHttpMethod     = http.MethodDelete
	eventsForPeriodHttpMethod = http.MethodGet
)

var serviceError = errors.New("service error")

type createEventResult struct {
	event types.Event
	err   error
}

type createEventInputData struct {
	createEventData          types.EventHandlerData
	createEventServiceReturn createEventResult
}

type createEventExpectedResult struct {
	createEventClientResult createEventResult
}

func TestCalendar_CreateEvent(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      createEventInputData
		expectedResult createEventExpectedResult
	}{
		{
			testCaseName: createEventTestCaseName,
			inputData: createEventInputData{
				createEventData: types.EventHandlerData{
					UserId: 0,
					Event: types.Event{
						Name: "event0",
						Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
					},
				},
				createEventServiceReturn: createEventResult{
					event: types.Event{
						Name: "event0",
						Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			expectedResult: createEventExpectedResult{
				createEventClientResult: createEventResult{
					event: types.Event{
						Name: "event0",
						Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
					},
				},
			},
		},
		{
			testCaseName: createEventServiceErrorTestCaseName,
			inputData: createEventInputData{
				createEventData: types.EventHandlerData{
					UserId: 0,
					Event: types.Event{
						Name: "event0",
						Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
					},
				},
				createEventServiceReturn: createEventResult{
					event: types.Event{},
					err:   serviceError,
				},
			},
			expectedResult: createEventExpectedResult{
				createEventClientResult: createEventResult{
					event: types.Event{},
					err:   serviceError,
				},
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			calendarServiceMock := &mocks.Service{}
			calendarServiceMock.On(
				"CreateEvent",
				mock.MatchedBy(func(ctx context.Context) bool {
					return true
				}),
				mock.MatchedBy(func(data types.EventHandlerData) bool {
					return data == testData.inputData.createEventData
				}),
			).Return(testData.inputData.createEventServiceReturn.event, testData.inputData.createEventServiceReturn.err)

			eventServer := getEventTestServer(calendarServiceMock)
			defer eventServer.Close()
			calendarClient := getCalendarClient(eventServer.URL)

			mainCtx := context.Background()
			ctx, cancel := context.WithCancel(mainCtx)
			defer cancel()
			event, err := calendarClient.CreateEvent(ctx, testData.inputData.createEventData)

			assert.Equal(t, testData.expectedResult.createEventClientResult.event, event)
			assert.Equal(t, testData.expectedResult.createEventClientResult.err, err)
		})
	}
}

type updateEventResult struct {
	event types.Event
	err   error
}

type updateEventInputData struct {
	updateEventData          types.UpdateEventHandlerData
	updateEventServiceReturn updateEventResult
}

type updateEventExpectedResult struct {
	updateEventClientResult updateEventResult
}

func TestCalendar_UpdateEvent(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      updateEventInputData
		expectedResult updateEventExpectedResult
	}{
		{
			testCaseName: updateEventTestCaseName,
			inputData: updateEventInputData{
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
				updateEventServiceReturn: updateEventResult{
					event: types.Event{
						Name: "event1",
						Date: time.Date(2022, time.January, 14, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			expectedResult: updateEventExpectedResult{
				updateEventClientResult: updateEventResult{
					event: types.Event{
						Name: "event1",
						Date: time.Date(2022, time.January, 14, 0, 0, 0, 0, time.UTC),
					},
				},
			},
		},
		{
			testCaseName: updateEventServiceErrorTestCaseName,
			inputData: updateEventInputData{
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
				updateEventServiceReturn: updateEventResult{
					event: types.Event{},
					err:   serviceError,
				},
			},
			expectedResult: updateEventExpectedResult{
				updateEventClientResult: updateEventResult{
					event: types.Event{},
					err:   serviceError,
				},
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			calendarServiceMock := &mocks.Service{}
			calendarServiceMock.On(
				"UpdateEvent",
				mock.MatchedBy(func(ctx context.Context) bool {
					return true
				}),
				mock.MatchedBy(func(data types.UpdateEventHandlerData) bool {
					return data == testData.inputData.updateEventData
				}),
			).Return(testData.inputData.updateEventServiceReturn.event, testData.inputData.updateEventServiceReturn.err)

			eventServer := getEventTestServer(calendarServiceMock)
			defer eventServer.Close()
			calendarClient := getCalendarClient(eventServer.URL)

			mainCtx := context.Background()
			ctx, cancel := context.WithCancel(mainCtx)
			defer cancel()
			event, err := calendarClient.UpdateEvent(ctx, testData.inputData.updateEventData)

			assert.Equal(t, testData.expectedResult.updateEventClientResult.event, event)
			assert.Equal(t, testData.expectedResult.updateEventClientResult.err, err)
		})
	}
}

type deleteEventResult struct {
	err error
}

type deleteEventInputData struct {
	deleteEventData          types.EventHandlerData
	deleteEventServiceReturn deleteEventResult
}

type deleteEventExpectedResult struct {
	deleteEventClientResult deleteEventResult
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
				deleteEventData: types.EventHandlerData{
					UserId: 0,
					Event: types.Event{
						Name: "event0",
						Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
					},
				},
				deleteEventServiceReturn: deleteEventResult{},
			},
			expectedResult: deleteEventExpectedResult{
				deleteEventClientResult: deleteEventResult{},
			},
		},
		{
			testCaseName: deleteEventServiceErrorTestCaseName,
			inputData: deleteEventInputData{
				deleteEventData: types.EventHandlerData{
					UserId: 0,
					Event: types.Event{
						Name: "event0",
						Date: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
					},
				},
				deleteEventServiceReturn: deleteEventResult{
					err: serviceError,
				},
			},
			expectedResult: deleteEventExpectedResult{
				deleteEventClientResult: deleteEventResult{
					err: serviceError,
				},
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			calendarServiceMock := &mocks.Service{}
			calendarServiceMock.On(
				"DeleteEvent",
				mock.MatchedBy(func(ctx context.Context) bool {
					return true
				}),
				mock.MatchedBy(func(data types.EventHandlerData) bool {
					return data == testData.inputData.deleteEventData
				}),
			).Return(testData.inputData.deleteEventServiceReturn.err)

			eventServer := getEventTestServer(calendarServiceMock)
			defer eventServer.Close()
			calendarClient := getCalendarClient(eventServer.URL)

			mainCtx := context.Background()
			ctx, cancel := context.WithCancel(mainCtx)
			defer cancel()
			err := calendarClient.DeleteEvent(ctx, testData.inputData.deleteEventData)

			assert.Equal(t, testData.expectedResult.deleteEventClientResult.err, err)
		})
	}
}

type eventsForPeriodResult struct {
	events []types.Event
	err    error
}

type eventsForPeriodInputData struct {
	eventsForPeriodData          types.DateIntervalHandlerData
	eventsForPeriodServiceReturn eventsForPeriodResult
}

type eventsForPeriodExpectedResult struct {
	eventsForPeriodClientResult eventsForPeriodResult
}

func TestCalendar_EventsForPeriod(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      eventsForPeriodInputData
		expectedResult eventsForPeriodExpectedResult
	}{
		{
			testCaseName: eventsForPeriodTestCaseName,
			inputData: eventsForPeriodInputData{
				eventsForPeriodData: types.DateIntervalHandlerData{
					UserId:    0,
					StartDate: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
					EndDate:   time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC),
				},
				eventsForPeriodServiceReturn: eventsForPeriodResult{
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
			expectedResult: eventsForPeriodExpectedResult{
				eventsForPeriodClientResult: eventsForPeriodResult{
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
			testCaseName: eventsForPeriodServiceErrorTestCaseName,
			inputData: eventsForPeriodInputData{
				eventsForPeriodData: types.DateIntervalHandlerData{
					UserId:    0,
					StartDate: time.Date(2022, time.January, 13, 0, 0, 0, 0, time.UTC),
					EndDate:   time.Date(2022, time.January, 15, 0, 0, 0, 0, time.UTC),
				},
				eventsForPeriodServiceReturn: eventsForPeriodResult{
					events: nil,
					err:    serviceError,
				},
			},
			expectedResult: eventsForPeriodExpectedResult{
				eventsForPeriodClientResult: eventsForPeriodResult{
					events: nil,
					err:    serviceError,
				},
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			calendarServiceMock := &mocks.Service{}
			calendarServiceMock.On(
				"EventsForPeriod",
				mock.MatchedBy(func(ctx context.Context) bool {
					return true
				}),
				mock.MatchedBy(func(data types.DateIntervalHandlerData) bool {
					return data == testData.inputData.eventsForPeriodData
				}),
			).Return(testData.inputData.eventsForPeriodServiceReturn.events, testData.inputData.eventsForPeriodServiceReturn.err)

			eventServer := getEventsForPeriodTestServer(calendarServiceMock)
			defer eventServer.Close()
			calendarClient := getCalendarClient(eventServer.URL)

			mainCtx := context.Background()
			ctx, cancel := context.WithCancel(mainCtx)
			defer cancel()
			events, err := calendarClient.EventsForPeriod(ctx, testData.inputData.eventsForPeriodData)

			assert.Equal(t, testData.expectedResult.eventsForPeriodClientResult.events, events)
			assert.Equal(t, testData.expectedResult.eventsForPeriodClientResult.err, err)
		})
	}
}

func getEventTestServer(service service.Service) *httptest.Server {
	createEventTransport := serverTransport.NewCreateEventTransport(dateFormat)
	updateEventTransport := serverTransport.NewUpdateEventTransport(dateFormat)
	deleteEventTransport := serverTransport.NewDeleteEventTransport(dateFormat)
	errorTransport := serverTransport.NewErrorTransport()

	return httptest.NewServer(handler.NewEventServer(
		createEventTransport,
		updateEventTransport,
		deleteEventTransport,
		service,
		errorTransport,
	))
}

func getEventsForPeriodTestServer(service service.Service) *httptest.Server {
	eventsForPeriodTransport := serverTransport.NewEventsForPeriodTransport(dateFormat)
	errorTransport := serverTransport.NewErrorTransport()

	return httptest.NewServer(handler.NewEventsForPeriodServer(
		eventsForPeriodTransport,
		service,
		errorTransport,
	))
}

func getCalendarClient(serverURL string) service.Service {
	errorClientTransport := clientTransport.NewErrorClientTransport()
	var parsedUrl *url.URL
	parsedUrl, _ = url.Parse(serverURL)
	createEventClientTransport := clientTransport.NewCreateEventClientTransport(
		parsedUrl,
		createEventHttpMethod,
		errorClientTransport,
		dateFormat,
	)
	updateEventClientTransport := clientTransport.NewUpdateEventClientTransport(
		parsedUrl,
		updateEventHttpMethod,
		errorClientTransport,
		dateFormat,
	)
	deleteEventClientTransport := clientTransport.NewDeleteEventClientTransport(
		parsedUrl,
		deleteEventHttpMethod,
		errorClientTransport,
		dateFormat,
	)
	parsedUrl, _ = url.Parse(serverURL)
	eventsForPeriodClientTransport := clientTransport.NewEventsForPeriodClientTransport(
		parsedUrl,
		eventsForPeriodHttpMethod,
		errorClientTransport,
		dateFormat,
	)
	return client.NewClient(
		createEventClientTransport,
		updateEventClientTransport,
		deleteEventClientTransport,
		eventsForPeriodClientTransport,
	)
}
