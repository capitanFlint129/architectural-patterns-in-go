package middleware

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

const (
	createEventRequestDurationMetricLabel     = "create_event"
	updateEventRequestDurationMetricLabel     = "update_event"
	deleteEventRequestDurationMetricLabel     = "delete_event"
	eventsForPeriodRequestDurationMetricLabel = "event_for_month"
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

func (r *requestDurationMiddleware) EventsForPeriod(ctx context.Context, data types.DateIntervalHandlerData) ([]types.Event, error) {
	start := time.Now()
	defer r.requestDurationMetric.WithLabelValues(eventsForPeriodRequestDurationMetricLabel).Observe(time.Since(start).Seconds())
	return r.service.EventsForPeriod(ctx, data)
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
