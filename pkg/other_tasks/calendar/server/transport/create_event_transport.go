package transport

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type createEventTransport struct {
	dateFormat string
}

func (c *createEventTransport) DecodeRequest(r *http.Request) (types.EventHandlerData, error) {
	return getEventHandlerDataFromRequest(r, c.dateFormat)
}

func (c *createEventTransport) EncodeResponse(w http.ResponseWriter, event types.Event) error {
	return encodeEventResponse(w, event, http.StatusCreated)
}

func NewCreateEventTransport(dateFormat string) CreateEventTransport {
	return &createEventTransport{
		dateFormat: dateFormat,
	}
}
