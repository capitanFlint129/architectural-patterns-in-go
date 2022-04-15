package transport

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type eventsForDayTransport struct {
	dateFormat string
}

func (c *eventsForDayTransport) DecodeRequest(r *http.Request) (types.DateHandlerData, error) {
	return getDateHandlerDataFromRequest(r, c.dateFormat)
}

func (c *eventsForDayTransport) EncodeResponse(w http.ResponseWriter, events []types.Event) error {
	return encodeEventsListResponse(w, events, http.StatusOK)
}

func NewEventsForDayTransport(dateFormat string) EventsForDayTransport {
	return &eventsForDayTransport{
		dateFormat: dateFormat,
	}
}
