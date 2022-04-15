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

type deleteEventClientTransport struct {
	url            *url.URL
	httpMethod     string
	errorTransport ErrorClientTransport
	dateFormat     string
}

func (c *deleteEventClientTransport) EncodeRequest(data types.EventHandlerData) (*http.Request, error) {
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

func (c *deleteEventClientTransport) DecodeResponse(r *http.Response) error {
	if r.StatusCode != http.StatusNoContent {
		return c.errorTransport.DecodeError(r)
	}

	var response types.EventResponse
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return err
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
