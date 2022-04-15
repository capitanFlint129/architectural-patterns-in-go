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

type updateEventClientTransport struct {
	url            *url.URL
	path           string
	httpMethod     string
	errorTransport ErrorClientTransport
	dateFormat     string
}

func (c *updateEventClientTransport) EncodeRequest(data types.EventHandlerData) (*http.Request, error) {
	params := url.Values{}
	params.Set("user_id", strconv.Itoa(data.UserId))
	params.Set("name", data.Event.Name)
	params.Set("date", data.Event.Date.Format(c.dateFormat))

	r, err := http.NewRequest(c.httpMethod, c.url.String(), strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return r, nil
}

func (c *updateEventClientTransport) DecodeResponse(r *http.Response) (types.Event, error) {
	if r.StatusCode != http.StatusOK {
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

func NewUpdateEventClientTransport(
	url *url.URL, path string, httpMethod string, errorTransport ErrorClientTransport, dateFormat string,
) UpdateEventClientTransport {
	return &updateEventClientTransport{
		url:            url,
		path:           path,
		httpMethod:     httpMethod,
		errorTransport: errorTransport,
		dateFormat:     dateFormat,
	}
}
