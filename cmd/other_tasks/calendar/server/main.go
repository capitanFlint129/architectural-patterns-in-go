package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/server/handler"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/server/middleware"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/server/service"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/server/transport"
)

// TODO обрабатывать обязательные и необязательные поля, их формат, допустимые значения и прочее

const (
	dateFormat = "2006-01-02"

	addr = ":8080"

	createEventPathPattern    = "/create_event"
	updateEventPathPattern    = "/update_event"
	deleteEventPathPattern    = "/delete_event"
	eventsForDayPathPattern   = "/event_for_day"
	eventsForWeekPathPattern  = "/event_for_week"
	eventsForMonthPathPattern = "/event_for_month"

	prometheusPathPattern = "/metrics"
)

func main() {
	logger := logrus.New()
	requestDurationMetric := promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "request_duration",
			Help:    "Duration of request metric",
			Buckets: prometheus.LinearBuckets(0, 1, 5),
		},
		[]string{"handler"},
	)

	calendarService := middleware.NewRequestDurationMiddleware(
		middleware.NewLoggingMiddleware(service.NewCalendar(), logger, dateFormat),
		requestDurationMetric,
	)
	mux := http.NewServeMux()
	createEventTransport := transport.NewCreateEventTransport(dateFormat)
	updateEventTransport := transport.NewUpdateEventTransport(dateFormat)
	deleteEventTransport := transport.NewDeleteEventTransport(dateFormat)
	eventsForDayTransport := transport.NewEventsForDayTransport(dateFormat)
	eventsForWeekTransport := transport.NewEventsForWeekTransport(dateFormat)
	eventsForMonthTransport := transport.NewEventsForMonthTransport(dateFormat)
	errorTransport := transport.NewErrorTransport()

	mux.Handle(createEventPathPattern, handler.NewCreateEventHandler(createEventTransport, calendarService, errorTransport))
	mux.Handle(updateEventPathPattern, handler.NewUpdateEventHandler(updateEventTransport, calendarService, errorTransport))
	mux.Handle(deleteEventPathPattern, handler.NewDeleteEventHandler(deleteEventTransport, calendarService, errorTransport))
	mux.Handle(eventsForDayPathPattern, handler.NewEventsForDayHandler(eventsForDayTransport, calendarService, errorTransport))
	mux.Handle(eventsForWeekPathPattern, handler.NewEventsForWeekHandler(eventsForWeekTransport, calendarService, errorTransport))
	mux.Handle(eventsForMonthPathPattern, handler.NewEventsForMonthHandler(eventsForMonthTransport, calendarService, errorTransport))
	mux.Handle(prometheusPathPattern, promhttp.Handler())
	calendarServer := http.Server{
		Addr:    addr,
		Handler: mux,
	}
	err := calendarServer.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
	}
}
