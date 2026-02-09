package tracer

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/bridges/otelzap"
	otelmetric "go.opentelemetry.io/otel/metric"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type TelemetryProvider interface {
	GetServiceName() string
	Info(args ...any)
	Error(args ...any)
	Fatal(args ...any)
	MeterInt64Histogram(metric Metric) (otelmetric.Int64Histogram, error)
	MeterInt64UpDownCounter(metric Metric) (otelmetric.Int64UpDownCounter, error)
	TraceStart(ctx context.Context, name string) (context.Context, oteltrace.Span)
	LogRequest() gin.HandlerFunc
	MeterRequestDuration() gin.HandlerFunc
	MeterRequestInFlight() gin.HandlerFunc
	Shutdown(ctx context.Context)
}

type Telemetry struct {
	lp          *sdklog.LoggerProvider
	mp          *metric.MeterProvider
	tp          *trace.TracerProvider
	log         *zap.SugaredLogger
	meter       otelmetric.Meter
	tracer      oteltrace.Tracer
	serviceName string
}

func NewTelemetry(ctx context.Context, serviceName, serviceVersion, endpoint string) (*Telemetry, error) {
	res := newResource(serviceName, serviceVersion)

	lp, err := newLoggerProvider(ctx, res, endpoint)
	if err != nil {
		return nil, err
	}

	logger := zap.New(
		// zapcore.NewTee(
		// 	zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
		otelzap.NewCore(serviceName, otelzap.WithLoggerProvider(lp)),
		// ),
	)

	mp, err := newMeterProvider(ctx, res, endpoint)
	if err != nil {
		return nil, err
	}
	meter := mp.Meter(serviceName)

	tp, err := newTraceProvider(ctx, res, endpoint)
	if err != nil {
		return nil, err
	}
	tracer := tp.Tracer(serviceName)

	return &Telemetry{
		lp:          lp,
		mp:          mp,
		tp:          tp,
		log:         logger.Sugar(),
		meter:       meter,
		tracer:      tracer,
		serviceName: serviceName,
	}, nil
}

func (t *Telemetry) GetServiceName() string {
	return t.serviceName
}

func (t *Telemetry) Info(args ...any) {
	t.log.Info(args...)
}

func (t *Telemetry) Error(args ...any) {
	t.log.Error(args...)
}

func (t *Telemetry) Fatal(args ...any) {
	t.log.Fatal(args...)
}

func (t *Telemetry) MeterInt64Histogram(metric Metric) (otelmetric.Int64Histogram, error) {
	histogram, err := t.meter.Int64Histogram(
		metric.Name,
		otelmetric.WithDescription(metric.Description),
		otelmetric.WithUnit(metric.Unit),
	)
	if err != nil {
		return nil, err
	}
	return histogram, nil
}

func (t *Telemetry) MeterInt64UpDownCounter(metric Metric) (otelmetric.Int64UpDownCounter, error) {
	counter, err := t.meter.Int64UpDownCounter(
		metric.Name,
		otelmetric.WithDescription(metric.Description),
		otelmetric.WithUnit(metric.Unit),
	)
	if err != nil {
		return nil, err
	}
	return counter, nil
}

func (t *Telemetry) TraceStart(ctx context.Context, name string) (context.Context, oteltrace.Span) {
	return t.tracer.Start(ctx, name)
}

func (t *Telemetry) Shutdown(ctx context.Context) {
	t.lp.Shutdown(ctx)
	t.mp.Shutdown(ctx)
	t.tp.Shutdown(ctx)
}
