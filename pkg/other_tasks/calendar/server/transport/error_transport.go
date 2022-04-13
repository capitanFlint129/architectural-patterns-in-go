package transport

import (
	"encoding/json"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type errorTransport struct {
}

func (c *errorTransport) EncodeError(w http.ResponseWriter, err error, statusCode int) {
	response := types.ErrorResponse{
		ErrorMsg: err.Error(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func NewErrorTransport() ErrorTransport {
	return &errorTransport{}
}
