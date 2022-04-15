package middleware

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
)

type service interface {
	CreateEvent(ctx context.Context, data types.HandlerEventData) (types.Event, error)
}
