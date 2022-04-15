package middleware

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
)

type service interface {
	CreateEvent(ctx context.Context, data types.EventHandlerData) (types.Event, error)
	UpdateEvent(ctx context.Context, data types.EventHandlerData) (types.Event, error)
	DeleteEvent(ctx context.Context, data types.EventHandlerData) error
	EventsForDay(ctx context.Context, data types.DateHandlerData) ([]types.Event, error)
	EventsForWeek(ctx context.Context, data types.DateHandlerData) ([]types.Event, error)
	EventsForMonth(ctx context.Context, data types.DateHandlerData) ([]types.Event, error)
}
