package middleware

import (
	"context"
	"github.com/capitanFlint129/architectural-patterns-in-go/pkg/other_tasks/calendar/types"
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"time"
)

type requestDurationMiddleware struct {
	service               service
	requestDurationMetric prometheus.Histogram
}

func (r *requestDurationMiddleware) CreateEvent(ctx context.Context, data types.HandlerEventData) (types.Event, error) {
	start := time.Now()
	createdEvent, err := r.service.CreateEvent(ctx, data)
	time.Sleep(time.Duration(rand.Int()%5) * time.Second)
	r.requestDurationMetric.Observe(time.Since(start).Seconds())
	return createdEvent, err
}

func NewRequestDurationMiddleware(service service, requestDurationMetric prometheus.Histogram) service {
	return &requestDurationMiddleware{
		service:               service,
		requestDurationMetric: requestDurationMetric,
	}
}
