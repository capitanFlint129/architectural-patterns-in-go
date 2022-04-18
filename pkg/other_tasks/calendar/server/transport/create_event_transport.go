package transport

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
	"strconv"
	"time"
)

type createEventTransport struct {
	dateFormat string
}

func (c *createEventTransport) DecodeRequest(r *http.Request) (types.EventHandlerData, error) {
	var (
		userId    int
		eventName string
		date      time.Time
		err       error
	)
	userId, err = strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		return types.EventHandlerData{}, err
	}
	eventName = r.FormValue("name")
	date, err = time.Parse(c.dateFormat, r.FormValue("date"))
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

func (c *createEventTransport) EncodeResponse(w http.ResponseWriter, event types.Event) error {
	return encodeEventResponse(w, event, http.StatusCreated)
}

func NewCreateEventTransport(dateFormat string) CreateEventTransport {
	return &createEventTransport{
		dateFormat: dateFormat,
	}
}
