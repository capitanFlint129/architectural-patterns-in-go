package handler

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type updateEventHandler struct {
	transport      updateEventTransport
	calendar       service
	errorTransport errorTransport
}

func (c *updateEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		updateEventData types.UpdateEventData
		event           types.Event
		err             error
	)
	if r.Method != http.MethodPost {
		c.errorTransport.EncodeError(w, err, http.StatusMethodNotAllowed)
		return
	}

	updateEventData, err = c.transport.DecodeRequest(r)
	if err != nil {
		c.errorTransport.EncodeError(w, err, http.StatusBadRequest)
		return
	}

	mainCtx := context.Background()
	ctx, cancel := context.WithCancel(mainCtx)
	defer cancel()
	event, err = c.calendar.UpdateEvent(ctx, updateEventData)
	if err != nil {
		c.errorTransport.EncodeError(w, err, http.StatusServiceUnavailable)
		return
	}

	if err = c.transport.EncodeResponse(w, event); err != nil {
		c.errorTransport.EncodeError(w, err, http.StatusInternalServerError)
		return
	}
}

func NewUpdateEventHandler(transport updateEventTransport, calendar service, errorTransport errorTransport) http.Handler {
	return &updateEventHandler{
		transport:      transport,
		calendar:       calendar,
		errorTransport: errorTransport,
	}
}
