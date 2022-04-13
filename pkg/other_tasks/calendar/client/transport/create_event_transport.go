package transport

import (
	"encoding/json"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type createEventClientTransport struct {
	url             *url.URL
	createEventPath string
	httpMethod      string
	errorTransport  ErrorClientTransport
	dateFormat      string
}

func (c *createEventClientTransport) EncodeRequest(createEventData types.CreateEventData) (*http.Request, error) {
	params := url.Values{}
	params.Set("user_id", strconv.Itoa(createEventData.UserId))
	params.Set("name", createEventData.Event.Name)
	params.Set("date", createEventData.Event.Date.Format(c.dateFormat))

	r, err := http.NewRequest(c.httpMethod, c.url.String(), strings.NewReader(params.Encode()))
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

	var response types.CreateEventResponse
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

func NewCreateEventClientTransport(url *url.URL, createEventPath string, httpMethod string, errorTransport ErrorClientTransport, dateFormat string) CreateEventClientTransport {
	return &createEventClientTransport{
		url:             url,
		createEventPath: createEventPath,
		httpMethod:      httpMethod,
		errorTransport:  errorTransport,
		dateFormat:      dateFormat,
	}
}
