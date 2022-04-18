package transport

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
)

type updateEventTransport struct {
	dateFormat string
}

func (c *updateEventTransport) DecodeRequest(r *http.Request) (types.UpdateEventHandlerData, error) {
	var (
		userId    int
		eventName string
		date      time.Time
		newDate   time.Time
		err       error
	)
	userId, err = strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		return types.UpdateEventHandlerData{}, err
	}
	eventName = r.FormValue("name")
	date, err = time.Parse(c.dateFormat, r.FormValue("date"))
	if err != nil {
		return types.UpdateEventHandlerData{}, err
	}
	newEventName := r.FormValue("new_name")
	newDate, err = time.Parse(c.dateFormat, r.FormValue("new_date"))
	if err != nil {
		return types.UpdateEventHandlerData{}, err
	}
	return types.UpdateEventHandlerData{
		UserId: userId,
		Event: types.Event{
			Name: eventName,
			Date: date,
		},
		NewEvent: types.Event{
			Name: newEventName,
			Date: newDate,
		},
	}, nil
}

func (c *updateEventTransport) EncodeResponse(w http.ResponseWriter, event types.Event) error {
	response := types.EventResponse{
		Result: event,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		return err
	}
	return nil
}

func NewUpdateEventTransport(dateFormat string) UpdateEventTransport {
	return &updateEventTransport{
		dateFormat: dateFormat,
	}
}
