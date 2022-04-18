package handler

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type eventsForPeriodServer struct {
	transport      eventsForPeriodTransport
	calendar       service
	errorTransport errorTransport
}

func (c *eventsForPeriodServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		data   types.DateIntervalHandlerData
		events []types.Event
		err    error
	)
	if r.Method != http.MethodGet {
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
	events, err = c.calendar.EventsForPeriod(ctx, data)
	if err != nil {
		c.errorTransport.EncodeError(w, err, http.StatusServiceUnavailable)
		return
	}

	if err = c.transport.EncodeResponse(w, events); err != nil {
		c.errorTransport.EncodeError(w, err, http.StatusInternalServerError)
		return
	}
}

func NewEventsForPeriodServer(transport eventsForPeriodTransport, calendar service, errorTransport errorTransport) http.Handler {
	return &eventsForPeriodServer{
		transport:      transport,
		calendar:       calendar,
		errorTransport: errorTransport,
	}
}
