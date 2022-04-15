package handler

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type updateEventServer struct {
	transport      updateEventTransport
	calendar       service
	errorTransport errorTransport
}

func (c *updateEventServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		data  types.EventHandlerData
		event types.Event
		err   error
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
	event, err = c.calendar.UpdateEvent(ctx, data)
	if err != nil {
		c.errorTransport.EncodeError(w, err, http.StatusServiceUnavailable)
		return
	}

	if err = c.transport.EncodeResponse(w, event); err != nil {
		c.errorTransport.EncodeError(w, err, http.StatusInternalServerError)
		return
	}
}

func NewUpdateEventServer(transport updateEventTransport, calendar service, errorTransport errorTransport) http.Handler {
	return &updateEventServer{
		transport:      transport,
		calendar:       calendar,
		errorTransport: errorTransport,
	}
}
