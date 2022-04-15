package transport

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type eventsForWeekTransport struct {
	dateFormat string
}

func (c *eventsForWeekTransport) DecodeRequest(r *http.Request) (types.DateHandlerData, error) {
	return getDateHandlerDataFromRequest(r, c.dateFormat)
}

func (c *eventsForWeekTransport) EncodeResponse(w http.ResponseWriter, events []types.Event) error {
	return encodeEventsListResponse(w, events, http.StatusOK)
}

func NewEventsForWeekTransport(dateFormat string) EventsForWeekTransport {
	return &eventsForWeekTransport{
		dateFormat: dateFormat,
	}
}
