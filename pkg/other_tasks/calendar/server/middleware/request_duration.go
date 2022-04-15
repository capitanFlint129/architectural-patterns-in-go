package middleware

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

const (
	createEventRequestDurationMetricLabel    = "create_event"
	updateEventRequestDurationMetricLabel    = "update_event"
	deleteEventRequestDurationMetricLabel    = "delete_event"
	eventsForDayRequestDurationMetricLabel   = "event_for_day"
	eventsForWeekRequestDurationMetricLabel  = "event_for_week"
	eventsForMonthRequestDurationMetricLabel = "event_for_month"
)

type requestDurationMiddleware struct {
	service               service
	requestDurationMetric *prometheus.HistogramVec
}

func (r *requestDurationMiddleware) CreateEvent(ctx context.Context, data types.EventHandlerData) (types.Event, error) {
	start := time.Now()
	defer r.requestDurationMetric.WithLabelValues(createEventRequestDurationMetricLabel).Observe(time.Since(start).Seconds())
	return r.service.CreateEvent(ctx, data)
}

func (r *requestDurationMiddleware) UpdateEvent(ctx context.Context, data types.UpdateEventHandlerData) (types.Event, error) {
	start := time.Now()
	defer r.requestDurationMetric.WithLabelValues(updateEventRequestDurationMetricLabel).Observe(time.Since(start).Seconds())
	return r.service.UpdateEvent(ctx, data)
}

func (r *requestDurationMiddleware) DeleteEvent(ctx context.Context, data types.EventHandlerData) error {
	start := time.Now()
	defer r.requestDurationMetric.WithLabelValues(deleteEventRequestDurationMetricLabel).Observe(time.Since(start).Seconds())
	return r.service.DeleteEvent(ctx, data)
}

func (r *requestDurationMiddleware) EventsForDay(ctx context.Context, data types.DateHandlerData) ([]types.Event, error) {
	start := time.Now()
	defer r.requestDurationMetric.WithLabelValues(eventsForDayRequestDurationMetricLabel).Observe(time.Since(start).Seconds())
	return r.service.EventsForDay(ctx, data)
}

func (r *requestDurationMiddleware) EventsForWeek(ctx context.Context, data types.DateHandlerData) ([]types.Event, error) {
	start := time.Now()
	defer r.requestDurationMetric.WithLabelValues(eventsForWeekRequestDurationMetricLabel).Observe(time.Since(start).Seconds())
	return r.service.EventsForWeek(ctx, data)
}

func (r *requestDurationMiddleware) EventsForMonth(ctx context.Context, data types.DateHandlerData) ([]types.Event, error) {
	start := time.Now()
	defer r.requestDurationMetric.WithLabelValues(eventsForMonthRequestDurationMetricLabel).Observe(time.Since(start).Seconds())
	return r.service.EventsForMonth(ctx, data)
}

func NewRequestDurationMiddleware(
	service service,
	requestDurationMetric *prometheus.HistogramVec,
) service {
	return &requestDurationMiddleware{
		service:               service,
		requestDurationMetric: requestDurationMetric,
	}
}
