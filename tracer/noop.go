package tracer

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	otelmetric "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type NoopTelemetry struct {
	serviceName string
}

func NewNoopTelemetry(serviceName string) *NoopTelemetry {
	return &NoopTelemetry{
		serviceName: serviceName,
	}
}

func (t *NoopTelemetry) GetServiceName() string {
	return t.serviceName
}

func (t *NoopTelemetry) Info(args ...any) {
}

func (t *NoopTelemetry) Error(args ...any) {
}

func (t *NoopTelemetry) Fatal(args ...any) {
	os.Exit(1)
}

func (t *NoopTelemetry) MeterInt64Histogram(metric Metric) (otelmetric.Int64Histogram, error) {
	return nil, nil
}

func (t *NoopTelemetry) MeterInt64UpDownCounter(metric Metric) (otelmetric.Int64UpDownCounter, error) {
	return nil, nil
}

func (t *NoopTelemetry) TraceStart(ctx context.Context, name string) (context.Context, oteltrace.Span) {
	return ctx, trace.SpanFromContext(ctx)
}

func (t *NoopTelemetry) LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func (t *NoopTelemetry) MeterRequestDuration() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func (t *NoopTelemetry) MeterRequestInFlight() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func (t *NoopTelemetry) Shutdown(ctx context.Context) {
}
