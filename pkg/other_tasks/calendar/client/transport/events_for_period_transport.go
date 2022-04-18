package transport

import (
	"encoding/json"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type eventsForPeriodClientTransport struct {
	url            *url.URL
	httpMethod     string
	errorTransport ErrorClientTransport
	dateFormat     string
}

func (c *eventsForPeriodClientTransport) EncodeRequest(data types.DateIntervalHandlerData) (*http.Request, error) {
	params := url.Values{}
	params.Set("user_id", strconv.Itoa(data.UserId))
	params.Set("start_date", data.StartDate.Format(c.dateFormat))
	params.Set("end_date", data.EndDate.Format(c.dateFormat))
	c.url.RawQuery = params.Encode()

	r, err := http.NewRequest(c.httpMethod, c.url.String(), nil)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return r, nil
}

func (c *eventsForPeriodClientTransport) DecodeResponse(r *http.Response) ([]types.Event, error) {
	if r.StatusCode != http.StatusOK {
		return nil, c.errorTransport.DecodeError(r)
	}
	var response types.EventsListResponse
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return nil, err
	}
	return response.Result, nil
}

func NewEventsForPeriodClientTransport(
	url *url.URL, httpMethod string, errorTransport ErrorClientTransport, dateFormat string,
) EventsForPeriodClientTransport {
	return &eventsForPeriodClientTransport{
		url:            url,
		httpMethod:     httpMethod,
		errorTransport: errorTransport,
		dateFormat:     dateFormat,
	}
}
