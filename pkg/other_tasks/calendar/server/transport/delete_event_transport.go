package transport

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type deleteEventTransport struct {
	dateFormat string
}

func (c *deleteEventTransport) DecodeRequest(r *http.Request) (types.EventHandlerData, error) {
	return getEventHandlerDataFromRequest(r, c.dateFormat)
}

func (c *deleteEventTransport) EncodeResponse(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func NewDeleteEventTransport(dateFormat string) DeleteEventTransport {
	return &deleteEventTransport{
		dateFormat: dateFormat,
	}
}
