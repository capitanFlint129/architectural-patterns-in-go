package middleware

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"github.com/sirupsen/logrus"
)

const (
	createEventHandlerName     = "create event"
	updateEventHandlerName     = "update event"
	deleteEventHandlerName     = "delete event"
	eventsForPeriodHandlerName = "events for period"
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

func (l *loggingMiddleware) UpdateEvent(ctx context.Context, data types.UpdateEventHandlerData) (types.Event, error) {
	l.logUpdateEventHandlerData(updateEventHandlerName, data)
	return l.service.UpdateEvent(ctx, data)
}

func (l *loggingMiddleware) DeleteEvent(ctx context.Context, data types.EventHandlerData) error {
	l.logEventHandlerData(deleteEventHandlerName, data)
	return l.service.DeleteEvent(ctx, data)
}

func (l *loggingMiddleware) EventsForPeriod(ctx context.Context, data types.DateIntervalHandlerData) ([]types.Event, error) {
	l.logDateIntervalHandlerData(eventsForPeriodHandlerName, data)
	return l.service.EventsForPeriod(ctx, data)
}

func (l *loggingMiddleware) logEventHandlerData(handlerName string, data types.EventHandlerData) {
	l.logger.WithFields(logrus.Fields{
		"handler_name": handlerName,
		"user_id":      data.UserId,
		"name":         data.Event.Name,
		"date":         data.Event.Date.Format(l.logDateFormat),
	}).Info()
}

func (l *loggingMiddleware) logUpdateEventHandlerData(handlerName string, data types.UpdateEventHandlerData) {
	l.logger.WithFields(logrus.Fields{
		"handler_name": handlerName,
		"user_id":      data.UserId,
		"name":         data.Event.Name,
		"date":         data.Event.Date.Format(l.logDateFormat),
		"new_name":     data.NewEvent.Name,
		"new_date":     data.NewEvent.Date.Format(l.logDateFormat),
	}).Info()
}

func (l *loggingMiddleware) logDateIntervalHandlerData(handlerName string, data types.DateIntervalHandlerData) {
	l.logger.WithFields(logrus.Fields{
		"handler_name": handlerName,
		"user_id":      data.UserId,
		"start_date":   data.StartDate.Format(l.logDateFormat),
		"end_date":     data.EndDate.Format(l.logDateFormat),
	}).Info()
}

func NewLoggingMiddleware(service service, logger *logrus.Logger, logDateFormat string) service {
	return &loggingMiddleware{
		service:       service,
		logger:        logger,
		logDateFormat: logDateFormat,
	}
}
