package transport

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type eventsForMonthTransport struct {
	dateFormat string
}

func (c *eventsForMonthTransport) DecodeRequest(r *http.Request) (types.DateHandlerData, error) {
	return getDateHandlerDataFromRequest(r, c.dateFormat)
}

func (c *eventsForMonthTransport) EncodeResponse(w http.ResponseWriter, events []types.Event) error {
	return encodeEventsListResponse(w, events, http.StatusOK)
}

func NewEventsForMonthTransport(dateFormat string) EventsForMonthTransport {
	return &eventsForMonthTransport{
		dateFormat: dateFormat,
	}
}
