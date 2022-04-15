package transport

import (
	"encoding/json"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type eventsForWeekClientTransport struct {
	url            *url.URL
	path           string
	httpMethod     string
	errorTransport ErrorClientTransport
	dateFormat     string
}

func (c *eventsForWeekClientTransport) EncodeRequest(data types.HandlerDateData) (*http.Request, error) {
	params := url.Values{}
	params.Set("user_id", strconv.Itoa(data.UserId))
	params.Set("date", data.Date.Format(c.dateFormat))
	c.url.RawQuery = params.Encode()

	r, err := http.NewRequest(c.httpMethod, c.url.String(), nil)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return r, nil
}

func (c *eventsForWeekClientTransport) DecodeResponse(r *http.Response) ([]types.Event, error) {
	if r.StatusCode != http.StatusOK {
		return nil, c.errorTransport.DecodeError(r)
	}

	var response types.EventResponse
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func NewEventsForWeekClientTransport(
	url *url.URL, path string, httpMethod string, errorTransport ErrorClientTransport, dateFormat string,
) EventsForWeekClientTransport {
	return &eventsForWeekClientTransport{
		url:            url,
		path:           path,
		httpMethod:     httpMethod,
		errorTransport: errorTransport,
		dateFormat:     dateFormat,
	}
}
