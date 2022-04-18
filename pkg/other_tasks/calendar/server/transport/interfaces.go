package transport

import (
	"net/http"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
)

type CreateEventTransport interface {
	DecodeRequest(r *http.Request) (types.EventHandlerData, error)
	EncodeResponse(w http.ResponseWriter, event types.Event) error
}

type UpdateEventTransport interface {
	DecodeRequest(r *http.Request) (types.UpdateEventHandlerData, error)
	EncodeResponse(w http.ResponseWriter, event types.Event) error
}

type DeleteEventTransport interface {
	DecodeRequest(r *http.Request) (types.EventHandlerData, error)
	EncodeResponse(w http.ResponseWriter) error
}

type EventsForPeriodTransport interface {
	DecodeRequest(r *http.Request) (types.DateIntervalHandlerData, error)
	EncodeResponse(w http.ResponseWriter, events []types.Event) error
}

type ErrorTransport interface {
	EncodeError(w http.ResponseWriter, err error, statusCode int)
}
