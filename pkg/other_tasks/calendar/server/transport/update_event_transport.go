package transport

import (
	"encoding/json"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type updateEventTransport struct {
	dateFormat string
}

func (c *updateEventTransport) DecodeRequest(r *http.Request) (types.UpdateEventHandlerData, error) {
	var (
		data types.UpdateEventHandlerData
		err  error
	)
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return types.UpdateEventHandlerData{}, err
	}
	if err != nil {
		return types.UpdateEventHandlerData{}, err
	}
	return data, nil
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
