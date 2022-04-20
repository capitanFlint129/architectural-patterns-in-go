package handler

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type eventServer struct {
	createEventTransport createEventTransport
	updateEventTransport updateEventTransport
	deleteEventTransport deleteEventTransport
	calendar             service
	errorTransport       errorTransport
}

func (e *eventServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		e.createEvent(w, r)
	case http.MethodPut:
		e.updateEvent(w, r)
	case http.MethodDelete:
		e.deleteEvent(w, r)
	default:
		e.errorTransport.EncodeError(w, types.ErrorMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}
}

func (e *eventServer) createEvent(w http.ResponseWriter, r *http.Request) {
	var (
		data  types.EventHandlerData
		event types.Event
		err   error
	)

	data, err = e.createEventTransport.DecodeRequest(r)
	if err != nil {
		e.errorTransport.EncodeError(w, err, http.StatusBadRequest)
		return
	}

	event, err = e.calendar.CreateEvent(r.Context(), data)
	if err != nil {
		e.errorTransport.EncodeError(w, err, http.StatusServiceUnavailable)
		return
	}

	if err = e.createEventTransport.EncodeResponse(w, event); err != nil {
		e.errorTransport.EncodeError(w, err, http.StatusInternalServerError)
		return
	}
}

func (e *eventServer) updateEvent(w http.ResponseWriter, r *http.Request) {
	var (
		data  types.UpdateEventHandlerData
		event types.Event
		err   error
	)

	data, err = e.updateEventTransport.DecodeRequest(r)
	if err != nil {
		e.errorTransport.EncodeError(w, err, http.StatusBadRequest)
		return
	}
	event, err = e.calendar.UpdateEvent(r.Context(), data)
	switch err {
	case nil:
	case types.ErrorEventNotFound:
		e.errorTransport.EncodeError(w, err, http.StatusNotFound)
		return
	default:
		e.errorTransport.EncodeError(w, err, http.StatusServiceUnavailable)
		return
	}

	if err = e.updateEventTransport.EncodeResponse(w, event); err != nil {
		e.errorTransport.EncodeError(w, err, http.StatusInternalServerError)
		return
	}
}

func (e *eventServer) deleteEvent(w http.ResponseWriter, r *http.Request) {
	var (
		data types.EventHandlerData
		err  error
	)

	data, err = e.deleteEventTransport.DecodeRequest(r)
	if err != nil {
		e.errorTransport.EncodeError(w, err, http.StatusBadRequest)
		return
	}
	err = e.calendar.DeleteEvent(r.Context(), data)
	switch err {
	case nil:
	case types.ErrorEventNotFound:
		e.errorTransport.EncodeError(w, err, http.StatusNotFound)
		return
	default:
		e.errorTransport.EncodeError(w, err, http.StatusServiceUnavailable)
		return
	}

	if err = e.deleteEventTransport.EncodeResponse(w); err != nil {
		e.errorTransport.EncodeError(w, err, http.StatusInternalServerError)
		return
	}
}

func NewEventServer(
	createEventTransport createEventTransport,
	updateEventTransport updateEventTransport,
	deleteEventTransport deleteEventTransport,
	calendar service,
	errorTransport errorTransport,
) http.Handler {
	return &eventServer{
		createEventTransport: createEventTransport,
		updateEventTransport: updateEventTransport,
		deleteEventTransport: deleteEventTransport,
		calendar:             calendar,
		errorTransport:       errorTransport,
	}
}
