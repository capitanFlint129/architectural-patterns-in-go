package transport

import (
	"bytes"
	"encoding/json"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"io/ioutil"
	"net/http"
	"net/url"
)

type createEventClientTransport struct {
	url            *url.URL
	httpMethod     string
	errorTransport ErrorClientTransport
	dateFormat     string
}

func (c *createEventClientTransport) EncodeRequest(data types.EventHandlerData) (*http.Request, error) {
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

func (c *createEventClientTransport) DecodeResponse(r *http.Response) (types.Event, error) {
	if r.StatusCode != http.StatusCreated {
		return types.Event{}, c.errorTransport.DecodeError(r)
	}

	var response types.EventResponse
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return types.Event{}, err
	}
	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return types.Event{}, err
	}
	return response.Result, nil
}

func NewCreateEventClientTransport(
	url *url.URL, httpMethod string, errorTransport ErrorClientTransport, dateFormat string,
) CreateEventClientTransport {
	return &createEventClientTransport{
		url:            url,
		httpMethod:     httpMethod,
		errorTransport: errorTransport,
		dateFormat:     dateFormat,
	}
}
