package monitoring

import (
	"context"

	"github.com/renanmedina/dcp-broadcaster/utils"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type TraceUnit struct {
	Span       trace.Span
	ctx        context.Context
	childSpans []trace.Span
}

func (tc TraceUnit) AddEvent(eventName string, options ...trace.EventOption) {
	tc.Span.AddEvent(eventName, options...)
}

func (tc TraceUnit) End() {
	if len(tc.childSpans) > 0 {
		for _, childSpan := range tc.childSpans {
			childSpan.End()
		}
	}

	tc.Span.End()
}

func (tc *TraceUnit) NewChildSpan(spanName string) *TraceUnit {
	newUnit := NewTrace(spanName, tc.ctx)
	tc.childSpans = append(tc.childSpans, newUnit.Span)
	return &newUnit
}

var (
	configs = utils.GetConfigs()
	logger  = utils.GetApplicationLogger()
	tracer  trace.Tracer
)

func InitTracer() *sdktrace.TracerProvider {
	resources, err := NewResources()
	if err != nil {
		logger.Error("Could not set open telemetry resources: %v", err)
	}

	exporter, err := NewOtelTraceExporter()
	if err != nil {
		logger.Error("Could not create open telemetry exporter: %v", err)
	}

	traceProvider := NewTracerProvider(exporter, resources)
	otel.SetTracerProvider(traceProvider)
	tracer = otel.Tracer("dcp-broadcaster-worker")

	return traceProvider
}

func NewTraceConsoleExporter() (*stdouttrace.Exporter, error) {
	return stdouttrace.New(stdouttrace.WithPrettyPrint())
}

func NewResources() (*resource.Resource, error) {
	return resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", configs.SERVICE_NAME),
			attribute.String("service.environment", configs.ENVIRONMENT),
			attribute.String("environment", configs.ENVIRONMENT),
			attribute.String("library.language", "go"),
		),
	)
}

func NewOtelTraceExporter() (*otlptrace.Exporter, error) {
	return otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			otlptracegrpc.WithInsecure(),
			otlptracegrpc.WithEndpoint(configs.OPEN_TELEMETRY_COLLECTOR_URL),
		),
	)
}

func NewTracerProvider(traceExporter sdktrace.SpanExporter, resources *resource.Resource) *sdktrace.TracerProvider {
	return sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(traceExporter),
		sdktrace.WithResource(resources),
	)
}

func NewTrace(traceName string, context context.Context) TraceUnit {
	newSpanContext, span := tracer.Start(context, traceName)
	return TraceUnit{
		span,
		newSpanContext,
		make([]trace.Span, 0),
	}
}

func NewAttributes(attrs map[string]string) trace.SpanStartEventOption {
	var mappedAttributes []attribute.KeyValue

	for attrKey, attrVal := range attrs {
		mappedAttributes = append(mappedAttributes, attribute.String(attrKey, attrVal))
	}

	return trace.WithAttributes(mappedAttributes...)
}

func ReportErrorFor(trace TraceUnit, err error) {
	trace.Span.RecordError(err)
	trace.Span.SetStatus(codes.Error, err.Error())
}
