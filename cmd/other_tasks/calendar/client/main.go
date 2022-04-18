package main

import (
	"context"
	"fmt"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/client/client"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/client/transport"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"time"
)

const (
	calendarUrl = "http://127.0.0.1:8080"
	dateFormat  = "2006-01-02"

	eventPath             = "/event"
	createEventHttpMethod = http.MethodPost
	updateEventHttpMethod = http.MethodPut
	deleteEventHttpMethod = http.MethodDelete

	eventsForPeriodPath       = "/events_for_period"
	eventsForPeriodHttpMethod = http.MethodGet
)

func main() {
	logger := logrus.New()
	errorTransport := transport.NewErrorClientTransport()

	var parsedUrl *url.URL
	parsedUrl, _ = url.Parse(calendarUrl + eventPath)
	createEventTransport := transport.NewCreateEventClientTransport(
		parsedUrl,
		createEventHttpMethod,
		errorTransport,
		dateFormat,
	)
	updateEventTransport := transport.NewUpdateEventClientTransport(
		parsedUrl,
		updateEventHttpMethod,
		errorTransport,
		dateFormat,
	)
	deleteEventTransport := transport.NewDeleteEventClientTransport(
		parsedUrl,
		deleteEventHttpMethod,
		errorTransport,
		dateFormat,
	)
	parsedUrl, _ = url.Parse(calendarUrl + eventsForPeriodPath)
	eventsForPeriodTransport := transport.NewEventsForPeriodClientTransport(
		parsedUrl,
		eventsForPeriodHttpMethod,
		errorTransport,
		dateFormat,
	)
	calendarClient := client.NewClient(
		createEventTransport,
		updateEventTransport,
		deleteEventTransport,
		eventsForPeriodTransport,
	)

	mainCtx := context.Background()
	ctx, cancel := context.WithCancel(mainCtx)
	defer cancel()
	//var event types.Event
	var err error
	date := time.Now()

	//// create event
	//event, err = calendarClient.CreateEvent(
	//	ctx,
	//	types.EventHandlerData{
	//		UserId: 0,
	//		Event: types.Event{
	//			Name: "event1",
	//			Date: date,
	//		},
	//	})
	//if err != nil {
	//	logger.Error(err)
	//} else {
	//	logger.Info(event)
	//}
	//
	//// update event
	//event, err = calendarClient.UpdateEvent(
	//	ctx,
	//	types.UpdateEventHandlerData{
	//		UserId: 0,
	//		Event: types.Event{
	//			Name: "event1",
	//			Date: date,
	//		},
	//		NewEvent: types.Event{
	//			Name: "event2",
	//			Date: date.AddDate(0, 0, -1),
	//		},
	//	})
	//if err != nil {
	//	logger.Error(err)
	//} else {
	//	logger.Info(event)
	//}
	//
	//// delete event
	//err = calendarClient.DeleteEvent(
	//	ctx,
	//	types.EventHandlerData{
	//		UserId: 0,
	//		Event: types.Event{
	//			Name: "event2",
	//			Date: date.AddDate(0, 0, -1),
	//		},
	//	})
	//if err != nil {
	//	logger.Error(err)
	//} else {
	//	logger.Info(event)
	//}

	var events []types.Event
	date = time.Date(2022, 1, 15, 0, 0, 0, 0, date.Location())

	// create event for each day create two events
	for i := 0; i < 30; i++ {
		for j := 0; j < 2; j++ {
			_, err = calendarClient.CreateEvent(
				ctx,
				types.EventHandlerData{
					UserId: 0,
					Event: types.Event{
						Name: fmt.Sprintf("event%d_%d", i, j),
						Date: date.AddDate(0, 0, i),
					},
				})
			if err != nil {
				logger.Error(err)
			}
		}
	}

	// events for day
	events, err = calendarClient.EventsForPeriod(
		ctx,
		types.DateIntervalHandlerData{
			UserId:    0,
			StartDate: date,
			EndDate:   date.AddDate(0, 0, 7),
		},
	)
	if err != nil {
		logger.Error(err)
	} else {
		fmt.Println(events)
	}
}
