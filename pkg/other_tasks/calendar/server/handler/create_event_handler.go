package handler

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type createEventHandler struct {
	transport      createEventTransport
	calendar       service
	errorTransport errorTransport
}

func (c *createEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		createEventData types.CreateEventData
		event           types.Event
		err             error
	)
	if r.Method != http.MethodPost {
		c.errorTransport.EncodeError(w, err, http.StatusMethodNotAllowed)
		return
	}

	createEventData, err = c.transport.DecodeRequest(r)
	if err != nil {
		c.errorTransport.EncodeError(w, err, http.StatusBadRequest)
		return
	}

	mainCtx := context.Background()
	ctx, cancel := context.WithCancel(mainCtx)
	defer cancel()
	event, err = c.calendar.CreateEvent(ctx, createEventData)
	if err != nil {
		c.errorTransport.EncodeError(w, err, http.StatusInternalServerError)
		return
	}

	if err = c.transport.EncodeResponse(w, event); err != nil {
		c.errorTransport.EncodeError(w, err, http.StatusInternalServerError)
		return
	}
}

func NewCreateEventHandler(transport createEventTransport, calendar service, errorTransport errorTransport) http.Handler {
	return &createEventHandler{
		transport:      transport,
		calendar:       calendar,
		errorTransport: errorTransport,
	}
}
