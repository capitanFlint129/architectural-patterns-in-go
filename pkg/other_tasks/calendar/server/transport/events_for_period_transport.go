package transport

import (
	"encoding/json"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type eventsForPeriodTransport struct {
	dateFormat string
}

func (c *eventsForPeriodTransport) DecodeRequest(r *http.Request) (types.DateIntervalHandlerData, error) {
	var (
		data types.DateIntervalHandlerData
		err  error
	)
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return types.DateIntervalHandlerData{}, err
	}
	if err != nil {
		return types.DateIntervalHandlerData{}, err
	}
	return data, nil
}

func (c *eventsForPeriodTransport) EncodeResponse(w http.ResponseWriter, events []types.Event) error {
	response := types.EventsListResponse{
		Result: events,
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

func NewEventsForPeriodTransport(dateFormat string) EventsForPeriodTransport {
	return &eventsForPeriodTransport{
		dateFormat: dateFormat,
	}
}
