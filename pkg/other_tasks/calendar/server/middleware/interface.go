package middleware

import (
	"context"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
)

type service interface {
	CreateEvent(ctx context.Context, data types.EventHandlerData) (types.Event, error)
	UpdateEvent(ctx context.Context, data types.UpdateEventHandlerData) (types.Event, error)
	DeleteEvent(ctx context.Context, data types.EventHandlerData) error
	EventsForPeriod(ctx context.Context, data types.DateIntervalHandlerData) ([]types.Event, error)
}
