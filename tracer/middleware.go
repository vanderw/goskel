package tracer

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/semconv/v1.13.0/httpconv"
)

func (t *Telemetry) LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		t.Info("request to ", c.Request.URL.Path)
		c.Next()
		t.Info("end of request to ", c.Request.URL.Path)
	}
}

func (t *Telemetry) MeterRequestDuration() gin.HandlerFunc {
	histogram, err := t.MeterInt64Histogram(MetricRequestDurationMillis)
	if err != nil { // handle error
		t.Fatal("failed to create histogram", err)
	}
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		duration := time.Since(startTime)

		histogram.Record(
			c.Request.Context(),
			int64(duration.Milliseconds()),
			metric.WithAttributes(
				httpconv.ServerRequest(t.GetServiceName(), c.Request)...,
			),
		)
	}
}

func (t *Telemetry) MeterRequestInFlight() gin.HandlerFunc {
	counter, err := t.MeterInt64UpDownCounter(MetricRequestsInFlight)
	if err != nil { // handle error
		t.Fatal("failed to create counter", err)
	}
	return func(c *gin.Context) {
		attrs := metric.WithAttributes(
			httpconv.ServerRequest(t.GetServiceName(), c.Request)...,
		)
		counter.Add(c.Request.Context(), 1, attrs)
		c.Next()
		counter.Add(c.Request.Context(), -1, attrs)
	}
}
