package transport

import (
	"encoding/json"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type deleteEventTransport struct {
	dateFormat string
}

func (c *deleteEventTransport) DecodeRequest(r *http.Request) (types.EventHandlerData, error) {
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

func (c *deleteEventTransport) EncodeResponse(w http.ResponseWriter) error {
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func NewDeleteEventTransport(dateFormat string) DeleteEventTransport {
	return &deleteEventTransport{
		dateFormat: dateFormat,
	}
}
