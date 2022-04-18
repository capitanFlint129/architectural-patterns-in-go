package client

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"net/http"
)

type createEventClientTransport interface {
	EncodeRequest(data types.EventHandlerData) (*http.Request, error)
	DecodeResponse(r *http.Response) (types.Event, error)
}

type updateEventClientTransport interface {
	EncodeRequest(data types.UpdateEventHandlerData) (*http.Request, error)
	DecodeResponse(r *http.Response) (types.Event, error)
}

type deleteEventClientTransport interface {
	EncodeRequest(data types.EventHandlerData) (*http.Request, error)
	DecodeResponse(r *http.Response) error
}

type eventsForPeriodClientTransport interface {
	EncodeRequest(data types.DateIntervalHandlerData) (*http.Request, error)
	DecodeResponse(r *http.Response) ([]types.Event, error)
}

type service interface {
	CreateEvent(ctx context.Context, data types.EventHandlerData) (types.Event, error)
	UpdateEvent(ctx context.Context, data types.UpdateEventHandlerData) (types.Event, error)
	DeleteEvent(ctx context.Context, data types.EventHandlerData) error
	EventsForPeriod(ctx context.Context, data types.DateIntervalHandlerData) ([]types.Event, error)
}

type client struct {
	createEventClientTransport     createEventClientTransport
	updateEventClientTransport     updateEventClientTransport
	deleteEventClientTransport     deleteEventClientTransport
	eventsForPeriodClientTransport eventsForPeriodClientTransport
}

func (c *client) CreateEvent(ctx context.Context, data types.EventHandlerData) (types.Event, error) {
	var (
		req   *http.Request
		resp  *http.Response
		event types.Event
		err   error
	)
	req, err = c.createEventClientTransport.EncodeRequest(data)
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

func (c *client) UpdateEvent(ctx context.Context, data types.UpdateEventHandlerData) (types.Event, error) {
	var (
		req   *http.Request
		resp  *http.Response
		event types.Event
		err   error
	)
	req, err = c.updateEventClientTransport.EncodeRequest(data)
	if err != nil {
		return types.Event{}, err
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return types.Event{}, err
	}
	event, err = c.updateEventClientTransport.DecodeResponse(resp)
	if err != nil {
		return types.Event{}, err
	}
	return event, nil
}

func (c *client) DeleteEvent(ctx context.Context, data types.EventHandlerData) error {
	var (
		req  *http.Request
		resp *http.Response
		err  error
	)
	req, err = c.deleteEventClientTransport.EncodeRequest(data)
	if err != nil {
		return err
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	err = c.deleteEventClientTransport.DecodeResponse(resp)
	if err != nil {
		return err
	}
	return nil
}

func (c *client) EventsForPeriod(ctx context.Context, data types.DateIntervalHandlerData) ([]types.Event, error) {
	var (
		req    *http.Request
		resp   *http.Response
		events []types.Event
		err    error
	)
	req, err = c.eventsForPeriodClientTransport.EncodeRequest(data)
	if err != nil {
		return nil, err
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	events, err = c.eventsForPeriodClientTransport.DecodeResponse(resp)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func NewClient(
	createEventClientTransport createEventClientTransport,
	updateEventClientTransport updateEventClientTransport,
	deleteEventClientTransport deleteEventClientTransport,
	eventsForPeriodClientTransport eventsForPeriodClientTransport,
) service {
	return &client{
		createEventClientTransport:     createEventClientTransport,
		updateEventClientTransport:     updateEventClientTransport,
		deleteEventClientTransport:     deleteEventClientTransport,
		eventsForPeriodClientTransport: eventsForPeriodClientTransport,
	}
}
