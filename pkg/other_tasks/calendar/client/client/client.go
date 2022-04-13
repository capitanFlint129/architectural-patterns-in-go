package client

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type createEventClientTransport interface {
	EncodeRequest(createEventData types.CreateEventData) (*http.Request, error)
	DecodeResponse(r *http.Response) (types.Event, error)
}

type service interface {
	CreateEvent(ctx context.Context, createEventData types.CreateEventData) (types.Event, error)
}

type client struct {
	createEventClientTransport createEventClientTransport
}

func (c *client) CreateEvent(ctx context.Context, createEventData types.CreateEventData) (types.Event, error) {
	var (
		req   *http.Request
		resp  *http.Response
		event types.Event
		err   error
	)
	req, err = c.createEventClientTransport.EncodeRequest(createEventData)
	if err != nil {
		return types.Event{}, err
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return types.Event{}, err
	}
	event, err = c.createEventClientTransport.DecodeResponse(resp)
	if err != nil {
		return types.Event{}, err
	}
	return event, nil
}

func NewClient(createEventClientTransport createEventClientTransport) service {
	return &client{
		createEventClientTransport: createEventClientTransport,
	}
}
