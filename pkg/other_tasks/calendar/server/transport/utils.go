package transport

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
)

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

func getDateIntervalHandlerDataFromRequest(r *http.Request, dateFormat string) (types.DateIntervalHandlerData, error) {
	var (
		userId    int
		startDate time.Time
		endDate   time.Time
		err       error
	)
	userId, err = strconv.Atoi(r.FormValue("user_id"))
	if err != nil {
		return types.DateIntervalHandlerData{}, err
	}
	startDate, err = time.Parse(dateFormat, r.FormValue("start_date"))
	if err != nil {
		return types.DateIntervalHandlerData{}, err
	}
	endDate, err = time.Parse(dateFormat, r.FormValue("end_date"))
	if err != nil {
		return types.DateIntervalHandlerData{}, err
	}
	return types.DateIntervalHandlerData{
		UserId:    userId,
		StartDate: startDate,
		EndDate:   endDate,
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
