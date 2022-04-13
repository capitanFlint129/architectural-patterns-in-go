package main

import (
	"context"
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

	createEventPath       = "/create_event"
	createEventHttpMethod = http.MethodPost
)

func main() {
	logger := logrus.New()
	createEventParsedUrl, _ := url.Parse(calendarUrl + createEventPath)
	errorTransport := transport.NewErrorClientTransport()
	createEventTransport := transport.NewCreateEventClientTransport(
		createEventParsedUrl,
		createEventPath,
		createEventHttpMethod,
		errorTransport,
		dateFormat,
	)
	calendarClient := client.NewClient(createEventTransport)

	mainCtx := context.Background()
	ctx, cancel := context.WithCancel(mainCtx)
	defer cancel()

	var event types.Event
	var err error
	date := time.Now()
	event, err = calendarClient.CreateEvent(
		ctx,
		types.CreateEventData{
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
}
