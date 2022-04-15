package service

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
)

type Service interface {
	CreateEvent(ctx context.Context, data types.HandlerEventData) (types.Event, error)
	UpdateEvent(ctx context.Context, data types.HandlerEventData) (types.Event, error)
	DeleteEvent(ctx context.Context, data types.HandlerEventData) error
	EventsForDay(ctx context.Context, data types.HandlerDateData) ([]types.Event, error)
	EventsForWeek(ctx context.Context, data types.HandlerDateData) ([]types.Event, error)
	EventsForMonth(ctx context.Context, data types.HandlerDateData) ([]types.Event, error)
}

type calendar struct {
	events map[int]map[string]types.Event
}

func (c *calendar) CreateEvent(ctx context.Context, data types.HandlerEventData) (types.Event, error) {
	c.events[data.UserId][data.Event.Name] = data.Event
	return data.Event, nil
}

func (c *calendar) UpdateEvent(ctx context.Context, data types.HandlerEventData) (types.Event, error) {
	_, ok := c.events[data.UserId][data.Event.Name]
	if !ok {
		return types.Event{}, types.ErrorEventNotFound
	}
	c.events[data.UserId][data.Event.Name] = data.Event
	return data.Event, nil
}

func (c *calendar) DeleteEvent(ctx context.Context, data types.HandlerEventData) error {
	_, ok := c.events[data.UserId][data.Event.Name]
	if !ok {
		return types.ErrorEventNotFound
	}
	delete(c.events[data.UserId], data.Event.Name)
	return nil
}

func (c *calendar) EventsForDay(ctx context.Context, data types.HandlerDateData) ([]types.Event, error) {
	eventsForDay := make([]types.Event, 0)
	for _, event := range c.events[data.UserId] {
		y1, m1, d1 := data.Date.Date()
		y2, m2, d2 := event.Date.Date()
		if y1 == y2 && m1 == m2 && d1 == d2 {
			eventsForDay = append(eventsForDay, event)
		}
	}
	return eventsForDay, nil
}

func (c *calendar) EventsForWeek(ctx context.Context, data types.HandlerDateData) ([]types.Event, error) {
	eventsForDay := make([]types.Event, 0)
	endOfWeek := data.Date.AddDate(0, 0, 7)
	for _, event := range c.events[data.UserId] {
		if event.Date.After(data.Date) && event.Date.Before(endOfWeek) {
			eventsForDay = append(eventsForDay, event)
		}
	}
	return eventsForDay, nil
}

func (c *calendar) EventsForMonth(ctx context.Context, data types.HandlerDateData) ([]types.Event, error) {
	eventsForMonth := make([]types.Event, 0)
	year, month, _ := data.Date.Date()
	for _, event := range c.events[data.UserId] {
		eventYear, eventMonth, _ := event.Date.Date()
		if eventYear == year && eventMonth == month {
			eventsForMonth = append(eventsForMonth, event)
		}
	}
	return eventsForMonth, nil

}

func NewCalendar() Service {
	return &calendar{
		events: make(map[int]map[string]types.Event),
	}
}
