package middleware

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"github.com/sirupsen/logrus"
	"time"
)

type loggingMiddleware struct {
	service       service
	logDateFormat string
	logger        *logrus.Logger
}

func (l *loggingMiddleware) CreateEvent(ctx context.Context, createEventData types.CreateEventData) (types.Event, error) {
	// TODO под время отдельный middleware
	// TODO prometheus - логгер - сервис
	start := time.Now()
	createdEvent, err := l.service.CreateEvent(ctx, createEventData)
	duration := time.Since(start).Microseconds()

	l.logger.WithFields(logrus.Fields{
		"user_id":  createEventData.UserId,
		"name":     createdEvent.Name,
		"date":     createdEvent.Date.Format(l.logDateFormat),
		"duration": duration,
	}).Info()

	return createdEvent, err
}

func NewLoggingMiddleware(service service, logger *logrus.Logger, logDateFormat string) service {
	return &loggingMiddleware{
		service:       service,
		logger:        logger,
		logDateFormat: logDateFormat,
	}
}
