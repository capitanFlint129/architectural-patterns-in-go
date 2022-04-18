package transport

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type eventsForPeriodTransport struct {
	dateFormat string
}

func (c *eventsForPeriodTransport) DecodeRequest(r *http.Request) (types.DateIntervalHandlerData, error) {
	return getDateIntervalHandlerDataFromRequest(r, c.dateFormat)
}

func (c *eventsForPeriodTransport) EncodeResponse(w http.ResponseWriter, events []types.Event) error {
	return encodeEventsListResponse(w, events, http.StatusOK)
}

func NewEventsForPeriodTransport(dateFormat string) EventsForPeriodTransport {
	return &eventsForPeriodTransport{
		dateFormat: dateFormat,
	}
}
