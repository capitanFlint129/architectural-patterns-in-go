package handler

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
	"time"
)

type eventServer struct {
	createEventTransport createEventTransport
	updateEventTransport updateEventTransport
	deleteEventTransport deleteEventTransport
	calendar             service
	errorTransport       errorTransport
}

func (c *eventServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error

	mainCtx := r.Context()
	ctx, cancel := context.WithTimeout(mainCtx, 10*time.Second)
	r = r.WithContext(ctx)
	defer cancel()

	switch r.Method {
	case http.MethodPost:
		var (
			data  types.EventHandlerData
			event types.Event
		)

		data, err = c.createEventTransport.DecodeRequest(r)
		if err != nil {
			c.errorTransport.EncodeError(w, err, http.StatusBadRequest)
			return
		}

		event, err = c.calendar.CreateEvent(r.Context(), data)
		if err != nil {
			c.errorTransport.EncodeError(w, err, http.StatusServiceUnavailable)
			return
		}

		if err = c.createEventTransport.EncodeResponse(w, event); err != nil {
			c.errorTransport.EncodeError(w, err, http.StatusInternalServerError)
			return
		}
	case http.MethodPut:
		var (
			data  types.UpdateEventHandlerData
			event types.Event
		)

		data, err = c.updateEventTransport.DecodeRequest(r)
		if err != nil {
			c.errorTransport.EncodeError(w, err, http.StatusBadRequest)
			return
		}
		event, err = c.calendar.UpdateEvent(r.Context(), data)
		switch err {
		case nil:
		case types.ErrorEventNotFound:
			c.errorTransport.EncodeError(w, err, http.StatusNotFound)
			return
		default:
			c.errorTransport.EncodeError(w, err, http.StatusServiceUnavailable)
			return
		}

		if err = c.updateEventTransport.EncodeResponse(w, event); err != nil {
			c.errorTransport.EncodeError(w, err, http.StatusInternalServerError)
			return
		}
	case http.MethodDelete:
		var (
			data types.EventHandlerData
		)

		data, err = c.deleteEventTransport.DecodeRequest(r)
		if err != nil {
			c.errorTransport.EncodeError(w, err, http.StatusBadRequest)
			return
		}
		err = c.calendar.DeleteEvent(r.Context(), data)
		switch err {
		case nil:
		case types.ErrorEventNotFound:
			c.errorTransport.EncodeError(w, err, http.StatusNotFound)
			return
		default:
			c.errorTransport.EncodeError(w, err, http.StatusServiceUnavailable)
			return
		}

		if err = c.deleteEventTransport.EncodeResponse(w); err != nil {
			c.errorTransport.EncodeError(w, err, http.StatusInternalServerError)
			return
		}
	default:
		c.errorTransport.EncodeError(w, types.ErrorMethodNotAllowed, http.StatusMethodNotAllowed)
		return
	}
}

func NewCreateEventServer(
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
