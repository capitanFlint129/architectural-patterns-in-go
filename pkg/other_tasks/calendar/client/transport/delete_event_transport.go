package transport

import (
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
	"net/url"
	"strconv"
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
	c.url.RawQuery = params.Encode()

	r, err := http.NewRequest(c.httpMethod, c.url.String(), nil)
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
