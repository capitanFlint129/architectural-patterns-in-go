package handler

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type createEventTransport interface {
	DecodeRequest(r *http.Request) (types.CreateEventData, error)
	EncodeResponse(w http.ResponseWriter, event types.Event) error
}

type errorTransport interface {
	EncodeError(w http.ResponseWriter, err error, statusCode int)
}

type service interface {
	CreateEvent(ctx context.Context, createEventData types.CreateEventData) (types.Event, error)
}
