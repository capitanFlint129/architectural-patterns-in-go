package transport

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type CreateEventClientTransport interface {
	EncodeRequest(createEventData types.CreateEventData) (*http.Request, error)
	DecodeResponse(r *http.Response) (types.Event, error)
}

type ErrorClientTransport interface {
	DecodeError(r *http.Response) error
}
