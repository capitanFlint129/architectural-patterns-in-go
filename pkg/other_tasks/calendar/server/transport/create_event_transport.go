package transport

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
)

type createEventTransport struct {
	dateFormat string
}

func (c *createEventTransport) DecodeRequest(r *http.Request) (types.CreateEventData, error) {
	var (
		userId    int
		eventName string
		date      time.Time
		err       error
	)
	userId, err = strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		return types.CreateEventData{}, err
	}
	eventName = r.FormValue("name")
	date, err = time.Parse(c.dateFormat, r.FormValue("date"))
	if err != nil {
		return types.CreateEventData{}, err
	}
	return types.CreateEventData{
		UserId: userId,
		Event: types.Event{
			Name: eventName,
			Date: date,
		},
	}, nil
}

func (c *createEventTransport) EncodeResponse(w http.ResponseWriter, event types.Event) error {
	response := types.CreateEventResponse{
		Result: event,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		return err
	}
	return nil
}

func NewCreateEventTransport(dateFormat string) CreateEventTransport {
	return &createEventTransport{
		dateFormat: dateFormat,
	}
}
