package middleware

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/server/types"
	"github.com/sirupsen/logrus"
	"time"
)

// TODO вынести формат куда-то
const dateFormat = "2006-01-02"

type service interface {
	CreateEvent(userId int, eventName string, eventDate time.Time) (types.Event, error)
}

type loggingMiddleware struct {
	service service
	logger  *logrus.Logger
}

func (l *loggingMiddleware) CreateEvent(userId int, eventName string, eventDate time.Time) (types.Event, error) {

	start := time.Now()
	event, err := l.service.CreateEvent(userId, eventName, eventDate)
	duration := time.Since(start).Microseconds()

	l.logger.WithFields(logrus.Fields{
		"user_id":  userId,
		"name":     eventName,
		"date":     eventDate.Format(dateFormat),
		"duration": duration,
	}).Info()

	return event, err
}

func NewLoggingMiddleware(service service, logger *logrus.Logger) Middleware {
	return &loggingMiddleware{
		service: service,
		logger:  logger,
	}
}
