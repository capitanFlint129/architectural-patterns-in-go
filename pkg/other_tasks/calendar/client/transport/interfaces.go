package transport

import (
	"net/http"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
)

type CreateEventClientTransport interface {
	EncodeRequest(data types.EventHandlerData) (*http.Request, error)
	DecodeResponse(r *http.Response) (types.Event, error)
}

type UpdateEventClientTransport interface {
	EncodeRequest(data types.EventHandlerData) (*http.Request, error)
	DecodeResponse(r *http.Response) (types.Event, error)
}

type DeleteEventClientTransport interface {
	EncodeRequest(data types.EventHandlerData) (*http.Request, error)
	DecodeResponse(r *http.Response) error
}

type EventsForDayClientTransport interface {
	EncodeRequest(data types.DateHandlerData) (*http.Request, error)
	DecodeResponse(r *http.Response) ([]types.Event, error)
}

type EventsForWeekClientTransport interface {
	EncodeRequest(data types.DateHandlerData) (*http.Request, error)
	DecodeResponse(r *http.Response) ([]types.Event, error)
}

type EventsForMonthClientTransport interface {
	EncodeRequest(data types.DateHandlerData) (*http.Request, error)
	DecodeResponse(r *http.Response) ([]types.Event, error)
}

type ErrorClientTransport interface {
	DecodeError(r *http.Response) error
}
