package handler

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type createEventTransport interface {
	DecodeRequest(r *http.Request) (types.EventHandlerData, error)
	EncodeResponse(w http.ResponseWriter, event types.Event) error
}

type updateEventTransport interface {
	DecodeRequest(r *http.Request) (types.EventHandlerData, error)
	EncodeResponse(w http.ResponseWriter, event types.Event) error
}

type deleteEventTransport interface {
	DecodeRequest(r *http.Request) (types.EventHandlerData, error)
	EncodeResponse(w http.ResponseWriter) error
}

type eventsForDayTransport interface {
	DecodeRequest(r *http.Request) (types.DateHandlerData, error)
	EncodeResponse(w http.ResponseWriter, events []types.Event) error
}

type eventsForWeekTransport interface {
	DecodeRequest(r *http.Request) (types.DateHandlerData, error)
	EncodeResponse(w http.ResponseWriter, events []types.Event) error
}

type eventsForMonthTransport interface {
	DecodeRequest(r *http.Request) (types.DateHandlerData, error)
	EncodeResponse(w http.ResponseWriter, events []types.Event) error
}

type errorTransport interface {
	EncodeError(w http.ResponseWriter, err error, statusCode int)
}

type service interface {
	CreateEvent(ctx context.Context, data types.EventHandlerData) (types.Event, error)
	UpdateEvent(ctx context.Context, data types.EventHandlerData) (types.Event, error)
	DeleteEvent(ctx context.Context, data types.EventHandlerData) error
	EventsForDay(ctx context.Context, data types.DateHandlerData) ([]types.Event, error)
	EventsForWeek(ctx context.Context, data types.DateHandlerData) ([]types.Event, error)
	EventsForMonth(ctx context.Context, data types.DateHandlerData) ([]types.Event, error)
}
