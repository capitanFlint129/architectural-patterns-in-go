package service

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
)

type Service interface {
	CreateEvent(ctx context.Context, createEventData types.CreateEventData) (types.Event, error)
}

type calendar struct {
	events map[int][]types.Event
}

func (c *calendar) CreateEvent(ctx context.Context, createEventData types.CreateEventData) (types.Event, error) {
	c.events[createEventData.UserId] = append(c.events[createEventData.UserId], createEventData.Event)
	return createEventData.Event, nil
}

func NewCalendar() Service {
	return &calendar{
		events: make(map[int][]types.Event),
	}
}
