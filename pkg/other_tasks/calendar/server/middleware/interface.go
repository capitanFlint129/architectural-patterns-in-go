package middleware

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/server/types"
	"time"
)

type Middleware interface {
	CreateEvent(userId int, eventName string, eventDate time.Time) (types.Event, error)
}
