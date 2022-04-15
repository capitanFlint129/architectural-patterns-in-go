package middleware

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"github.com/sirupsen/logrus"
)

const (
	createEventHandlerName    = "create event"
	updateEventHandlerName    = "update event"
	deleteEventHandlerName    = "delete event"
	eventsForDayHandlerName   = "events for day"
	eventsForWeekHandlerName  = "events for week"
	eventsForMonthHandlerName = "events for month"
)

type loggingMiddleware struct {
	service       service
	logDateFormat string
	logger        *logrus.Logger
}

func (l *loggingMiddleware) CreateEvent(ctx context.Context, data types.EventHandlerData) (types.Event, error) {
	l.logEventHandlerData(createEventHandlerName, data)
	return l.service.CreateEvent(ctx, data)
}

func (l *loggingMiddleware) UpdateEvent(ctx context.Context, data types.EventHandlerData) (types.Event, error) {
	l.logEventHandlerData(updateEventHandlerName, data)
	return l.service.UpdateEvent(ctx, data)
}

func (l *loggingMiddleware) DeleteEvent(ctx context.Context, data types.EventHandlerData) error {
	l.logEventHandlerData(deleteEventHandlerName, data)
	return l.service.DeleteEvent(ctx, data)
}

func (l *loggingMiddleware) EventsForDay(ctx context.Context, data types.DateHandlerData) ([]types.Event, error) {
	l.logDateHandlerData(eventsForDayHandlerName, data)
	return l.service.EventsForDay(ctx, data)
}

func (l *loggingMiddleware) EventsForWeek(ctx context.Context, data types.DateHandlerData) ([]types.Event, error) {
	l.logDateHandlerData(eventsForWeekHandlerName, data)
	return l.service.EventsForWeek(ctx, data)
}

func (l *loggingMiddleware) EventsForMonth(ctx context.Context, data types.DateHandlerData) ([]types.Event, error) {
	l.logDateHandlerData(eventsForMonthHandlerName, data)
	return l.service.EventsForMonth(ctx, data)
}

func (l *loggingMiddleware) logEventHandlerData(handlerName string, data types.EventHandlerData) {
	l.logger.WithFields(logrus.Fields{
		"handler_name": handlerName,
		"user_id":      data.UserId,
		"name":         data.Event.Name,
		"date":         data.Event.Date.Format(l.logDateFormat),
	}).Info()
}

func (l *loggingMiddleware) logDateHandlerData(handlerName string, data types.DateHandlerData) {
	l.logger.WithFields(logrus.Fields{
		"handler_name": handlerName,
		"user_id":      data.UserId,
		"date":         data.Date.Format(l.logDateFormat),
	}).Info()
}

func NewLoggingMiddleware(service service, logger *logrus.Logger, logDateFormat string) service {
	return &loggingMiddleware{
		service:       service,
		logger:        logger,
		logDateFormat: logDateFormat,
	}
}
