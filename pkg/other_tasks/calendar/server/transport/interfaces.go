package transport

import (
	"net/http"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
)

type CreateEventTransport interface {
	DecodeRequest(r *http.Request) (types.CreateEventData, error)
	EncodeResponse(w http.ResponseWriter, event types.Event) error
}

type ErrorTransport interface {
	EncodeError(w http.ResponseWriter, err error, statusCode int)
}
