package handler

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type deleteEventHandler struct {
	transport      deleteEventTransport
	calendar       service
	errorTransport errorTransport
}

func (c *deleteEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		data types.EventHandlerData
		err  error
	)
	if r.Method != http.MethodPost {
		c.errorTransport.EncodeError(w, err, http.StatusMethodNotAllowed)
		return
	}

	data, err = c.transport.DecodeRequest(r)
	if err != nil {
		c.errorTransport.EncodeError(w, err, http.StatusBadRequest)
		return
	}

	mainCtx := context.Background()
	ctx, cancel := context.WithCancel(mainCtx)
	defer cancel()
	err = c.calendar.DeleteEvent(ctx, data)
	if err != nil {
		c.errorTransport.EncodeError(w, err, http.StatusServiceUnavailable)
		return
	}

	if err = c.transport.EncodeResponse(w); err != nil {
		c.errorTransport.EncodeError(w, err, http.StatusInternalServerError)
		return
	}
}

func NewDeleteEventHandler(transport deleteEventTransport, calendar service, errorTransport errorTransport) http.Handler {
	return &deleteEventHandler{
		transport:      transport,
		calendar:       calendar,
		errorTransport: errorTransport,
	}
}
