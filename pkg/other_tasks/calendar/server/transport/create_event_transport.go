package transport

import (
	"encoding/json"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type createEventTransport struct {
	dateFormat string
}

func (c *createEventTransport) DecodeRequest(r *http.Request) (types.EventHandlerData, error) {
	var (
		data types.EventHandlerData
		err  error
	)
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return types.EventHandlerData{}, err
	}
	if err != nil {
		return types.EventHandlerData{}, err
	}
	return data, nil
}

func (c *createEventTransport) EncodeResponse(w http.ResponseWriter, event types.Event) error {
	response := types.EventResponse{
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
