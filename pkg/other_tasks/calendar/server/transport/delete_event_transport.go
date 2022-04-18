package transport

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
	"strconv"
	"time"
)

type deleteEventTransport struct {
	dateFormat string
}

func (c *deleteEventTransport) DecodeRequest(r *http.Request) (types.EventHandlerData, error) {
	var (
		userId    int
		eventName string
		date      time.Time
		err       error
	)
	userId, err = strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		return types.EventHandlerData{}, err
	}
	eventName = r.URL.Query().Get("name")
	date, err = time.Parse(c.dateFormat, r.URL.Query().Get("date"))
	if err != nil {
		return types.EventHandlerData{}, err
	}
	return types.EventHandlerData{
		UserId: userId,
		Event: types.Event{
			Name: eventName,
			Date: date,
		},
	}, nil
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
