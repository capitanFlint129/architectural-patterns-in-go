package main

import (
	"context"
	"fmt"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"time"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/client/client"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/client/transport"
)

const (
	calendarUrl = "http://127.0.0.1:8080"
	dateFormat  = "2006-01-02"

	createEventPath          = "/create_event"
	createEventHttpMethod    = http.MethodPost
	updateEventPath          = "/update_event"
	updateEventHttpMethod    = http.MethodPost
	deleteEventPath          = "/delete_event"
	deleteEventHttpMethod    = http.MethodPost
	eventsForDayPath         = "/events_for_day"
	eventsForDayHttpMethod   = http.MethodGet
	eventsForWeekPath        = "/events_for_week"
	eventsForWeekHttpMethod  = http.MethodGet
	eventsForMonthPath       = "/events_for_month"
	eventsForMonthHttpMethod = http.MethodGet
)

func main() {
	errorTransport := transport.NewErrorClientTransport()
	logger := logrus.New()

	var parsedUrl *url.URL
	parsedUrl, _ = url.Parse(calendarUrl + createEventPath)
	createEventTransport := transport.NewCreateEventClientTransport(
		parsedUrl,
		createEventHttpMethod,
		errorTransport,
		dateFormat,
	)
	parsedUrl, _ = url.Parse(calendarUrl + updateEventPath)
	updateEventTransport := transport.NewUpdateEventClientTransport(
		parsedUrl,
		updateEventHttpMethod,
		errorTransport,
		dateFormat,
	)
	parsedUrl, _ = url.Parse(calendarUrl + deleteEventPath)
	deleteEventTransport := transport.NewDeleteEventClientTransport(
		parsedUrl,
		deleteEventHttpMethod,
		errorTransport,
		dateFormat,
	)
	parsedUrl, _ = url.Parse(calendarUrl + eventsForDayPath)
	eventsForDayTransport := transport.NewEventsForDayClientTransport(
		parsedUrl,
		eventsForDayHttpMethod,
		errorTransport,
		dateFormat,
	)
	parsedUrl, _ = url.Parse(calendarUrl + eventsForWeekPath)
	eventsForWeekTransport := transport.NewEventsForWeekClientTransport(
		parsedUrl,
		eventsForWeekHttpMethod,
		errorTransport,
		dateFormat,
	)
	parsedUrl, _ = url.Parse(calendarUrl + eventsForMonthPath)
	eventsForMonthTransport := transport.NewEventsForMonthClientTransport(
		parsedUrl,
		eventsForMonthHttpMethod,
		errorTransport,
		dateFormat,
	)
	calendarClient := client.NewClient(
		createEventTransport,
		updateEventTransport,
		deleteEventTransport,
		eventsForDayTransport,
		eventsForWeekTransport,
		eventsForMonthTransport,
	)

	mainCtx := context.Background()
	ctx, cancel := context.WithCancel(mainCtx)
	defer cancel()

	// create event
	var event types.Event
	var err error
	date := time.Now()
	event, err = calendarClient.CreateEvent(
		ctx,
		types.EventHandlerData{
			UserId: 0,
			Event: types.Event{
				Name: "event1",
				Date: date,
			},
		})
	if err != nil {
		logger.Error(err)
	} else {
		logger.Info(event)
	}

	// update event
	event, err = calendarClient.UpdateEvent(
		ctx,
		types.EventHandlerData{
			UserId: 0,
			Event: types.Event{
				Name: "event2",
				Date: date.AddDate(0, 0, -1),
			},
		})
	if err != nil {
		logger.Error(err)
	} else {
		logger.Info(event)
	}

	// delete event
	err = calendarClient.DeleteEvent(
		ctx,
		types.EventHandlerData{
			UserId: 0,
			Event: types.Event{
				Name: "event2",
				Date: date.AddDate(0, 0, -1),
			},
		})
	if err != nil {
		logger.Error(err)
	} else {
		logger.Info(event)
	}

	// prepare events for events getters
	const (
		testEventsInDay   = 2
		testEventsInWeek  = 14
		testEventsInMonth = 34
	)
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
	events, err = calendarClient.EventsForDay(
		ctx,
		types.DateHandlerData{
			UserId: 0,
			Date:   date,
		},
	)
	if err != nil {
		logger.Error(err)
	} else if len(events) == testEventsInDay {
		logger.Info("events for day works fine!")
	} else {
		logger.Errorf("Wrong events number, got %d instead of %d", len(events), testEventsInDay)
	}

	// events for week
	events, err = calendarClient.EventsForDay(
		ctx,
		types.DateHandlerData{
			UserId: 0,
			Date:   date,
		},
	)
	if err != nil {
		logger.Error(err)
	} else if len(events) == testEventsInWeek {
		logger.Info("events for week works fine!")
	} else {
		logger.Errorf("Wrong events number, got %d instead of %d", len(events), testEventsInWeek)
	}

	// events for month
	events, err = calendarClient.EventsForMonth(
		ctx,
		types.DateHandlerData{
			UserId: 0,
			Date:   date,
		},
	)
	if err != nil {
		logger.Error(err)
	} else if len(events) == testEventsInMonth {
		logger.Info("events for month works fine!")
	} else {
		logger.Errorf("Wrong events number, got %d instead of %d", len(events), testEventsInMonth)
	}
}
