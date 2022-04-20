package transport

import (
	"bytes"
	"encoding/json"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
	"net/url"
)

type deleteEventClientTransport struct {
	url            *url.URL
	httpMethod     string
	errorTransport ErrorClientTransport
	dateFormat     string
}

func (c *deleteEventClientTransport) EncodeRequest(data types.EventHandlerData) (*http.Request, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequest(c.httpMethod, c.url.String(), bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return r, nil

}

func (c *deleteEventClientTransport) DecodeResponse(r *http.Response) error {
	if r.StatusCode != http.StatusNoContent {
		return c.errorTransport.DecodeError(r)
	}
	return nil
}

func NewDeleteEventClientTransport(
	url *url.URL, httpMethod string, errorTransport ErrorClientTransport, dateFormat string,
) DeleteEventClientTransport {
	return &deleteEventClientTransport{
		url:            url,
		httpMethod:     httpMethod,
		errorTransport: errorTransport,
		dateFormat:     dateFormat,
	}
}
