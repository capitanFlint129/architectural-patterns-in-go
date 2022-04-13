package main

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/server/handler"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/server/middleware"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/server/service"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/server/transport"
)

// TODO обрабатывать обязательные и необязательные поля, их формат, допустимые значения и прочее

const (
	dateFormat = "2006-01-02"

	addr = ":8080"

	createEventPathPattern = "/create_event"
)

func main() {
	logger := logrus.New()
	calendarService := middleware.NewLoggingMiddleware(service.NewCalendar(), logger, dateFormat)
	mux := http.NewServeMux()
	createEventTransport := transport.NewCreateEventTransport(dateFormat)
	errorTransport := transport.NewErrorTransport()

	mux.Handle(createEventPathPattern, handler.NewCreateEventHandler(createEventTransport, calendarService, errorTransport))
	calendarServer := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	err := calendarServer.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
	}
}
