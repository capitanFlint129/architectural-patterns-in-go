package transport

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
)

type errorClientTransport struct {
}

func (e *errorClientTransport) DecodeError(r *http.Response) error {
	var response types.ErrorResponse
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return err
	}
	return errors.New(response.ErrorMsg)
}

func NewErrorClientTransport() ErrorClientTransport {
	return &errorClientTransport{}
}
