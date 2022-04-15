package transport

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
)

func getEventHandlerDataFromRequest(r *http.Request, dateFormat string) (types.EventHandlerData, error) {
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
	date, err = time.Parse(dateFormat, r.FormValue("date"))
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

func encodeEventResponse(w http.ResponseWriter, event types.Event, statusCode int) error {
	response := types.EventResponse{
		Result: event,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return err
	}
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		return err
	}
	return nil
}

func getDateHandlerDataFromRequest(r *http.Request, dateFormat string) (types.DateHandlerData, error) {
	var (
		userId int
		date   time.Time
		err    error
	)
	userId, err = strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		return types.DateHandlerData{}, err
	}
	date, err = time.Parse(dateFormat, r.FormValue("date"))
	if err != nil {
		return types.DateHandlerData{}, err
	}
	return types.DateHandlerData{
		UserId: userId,
		Date:   date,
	}, nil
}

func encodeEventsListResponse(w http.ResponseWriter, events []types.Event, statusCode int) error {
	response := types.EventsListResponse{
		Result: events,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return err
	}
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		return err
	}
	return nil
}
