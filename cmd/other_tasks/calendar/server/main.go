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

const (
	dateFormat = "2006-01-02"

	addr = ":8080"

	eventPathPattern           = "/event"
	eventsForPeriodPathPattern = "/events_for_period"
	prometheusPathPattern      = "/metrics"
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
	eventsForPeriodTransport := transport.NewEventsForPeriodTransport(dateFormat)
	errorTransport := transport.NewErrorTransport()

	mux.Handle(eventPathPattern, handler.NewCreateEventServer(
		createEventTransport,
		updateEventTransport,
		deleteEventTransport,
		calendarService,
		errorTransport,
	))
	mux.Handle(eventsForPeriodPathPattern, handler.NewEventsForPeriodServer(eventsForPeriodTransport, calendarService, errorTransport))
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
