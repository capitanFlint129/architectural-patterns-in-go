package service

import (
	"context"
	"sort"

	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
)

type Service interface {
	CreateEvent(ctx context.Context, data types.EventHandlerData) (types.Event, error)
	UpdateEvent(ctx context.Context, data types.UpdateEventHandlerData) (types.Event, error)
	DeleteEvent(ctx context.Context, data types.EventHandlerData) error
	EventsForPeriod(ctx context.Context, data types.DateIntervalHandlerData) ([]types.Event, error)
}

type calendar struct {
	events map[int]map[string]types.Event
}

func (c *calendar) CreateEvent(ctx context.Context, data types.EventHandlerData) (types.Event, error) {
	_, ok := c.events[data.UserId]
	if !ok {
		c.events[data.UserId] = make(map[string]types.Event)
	}
	_, ok = c.events[data.UserId][data.Event.Name]
	if ok {
		return types.Event{}, types.ErrorEventAlreadyExists
	}
	c.events[data.UserId][data.Event.Name] = data.Event
	return data.Event, nil
}

func (c *calendar) UpdateEvent(ctx context.Context, data types.UpdateEventHandlerData) (types.Event, error) {
	_, ok := c.events[data.UserId][data.Event.Name]
	if !ok {
		return types.Event{}, types.ErrorEventNotFound
	}
	if data.Event.Name != data.NewEvent.Name {
		delete(c.events[data.UserId], data.Event.Name)
	}
	c.events[data.UserId][data.NewEvent.Name] = data.NewEvent
	return data.NewEvent, nil
}

func (c *calendar) DeleteEvent(ctx context.Context, data types.EventHandlerData) error {
	_, ok := c.events[data.UserId][data.Event.Name]
	if !ok {
		return types.ErrorEventNotFound
	}
	delete(c.events[data.UserId], data.Event.Name)
	return nil
}

func (c *calendar) EventsForPeriod(ctx context.Context, data types.DateIntervalHandlerData) ([]types.Event, error) {
	eventsForPeriod := make([]types.Event, 0)
	for _, event := range c.events[data.UserId] {
		if event.Date.After(data.StartDate) && event.Date.Before(data.EndDate) {
			eventsForPeriod = append(eventsForPeriod, event)
		}
	}
	sort.SliceStable(eventsForPeriod, func(i, j int) bool {
		return eventsForPeriod[i].Date.Before(eventsForPeriod[j].Date)
	})
	return eventsForPeriod, nil
}

func NewCalendar() Service {
	return &calendar{
		events: make(map[int]map[string]types.Event),
	}
}
