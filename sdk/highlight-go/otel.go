package highlight

import (
	"context"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/samber/lo"

	"github.com/pkg/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

const OTLPDefaultEndpoint = "https://otel.highlight.io:4318"

const ErrorURLAttribute = "URL"

const DeprecatedProjectIDAttribute = "highlight_project_id"
const DeprecatedSessionIDAttribute = "highlight_session_id"
const DeprecatedRequestIDAttribute = "highlight_trace_id"
const DeprecatedSourceAttribute = "Source"
const ProjectIDAttribute = "highlight.project_id"
const SessionIDAttribute = "highlight.session_id"
const RequestIDAttribute = "highlight.trace_id"
const SourceAttribute = "highlight.source"
const TraceTypeAttribute = "highlight.type"
const TraceKeyAttribute = "highlight.key"

const LogEvent = "log"
const LogSeverityAttribute = "log.severity"
const LogMessageAttribute = "log.message"

const MetricEvent = "metric"
const MetricEventName = "metric.name"
const MetricEventValue = "metric.value"

type TraceType string

const TraceTypeNetworkRequest TraceType = "http.request"
const TraceTypeHighlightInternal TraceType = "highlight.internal"
const TraceTypePhoneHome TraceType = "highlight.phonehome"

type OTLP struct {
	tracerProvider *sdktrace.TracerProvider
}

type ErrorWithStack interface {
	Error() string
	StackTrace() errors.StackTrace
}

type highlightSampler struct {
	traceIDUpperBounds map[trace.SpanKind]uint64
	description        string
}

func (ts highlightSampler) ShouldSample(p sdktrace.SamplingParameters) sdktrace.SamplingResult {
	psc := trace.SpanContextFromContext(p.ParentContext)
	x := binary.BigEndian.Uint64(p.TraceID[8:16]) >> 1
	bound, ok := ts.traceIDUpperBounds[p.Kind]
	if !ok {
		bound = ts.traceIDUpperBounds[trace.SpanKindUnspecified]
	}
	if x < bound {
		return sdktrace.SamplingResult{
			Decision:   sdktrace.RecordAndSample,
			Tracestate: psc.TraceState(),
		}
	}
	return sdktrace.SamplingResult{
		Decision:   sdktrace.Drop,
		Tracestate: psc.TraceState(),
	}
}

func (ts highlightSampler) Description() string {
	return ts.description
}

// creates a per-span-kind sampler that samples each kind at a provided fraction.
func getSampler() highlightSampler {
	return highlightSampler{
		description: fmt.Sprintf("TraceIDRatioBased{%+v}", conf.samplingRateMap),
		traceIDUpperBounds: lo.MapEntries(conf.samplingRateMap, func(key trace.SpanKind, value float64) (trace.SpanKind, uint64) {
			return key, uint64(value * (1 << 63))
		}),
	}
}

var (
	tracer = otel.GetTracerProvider().Tracer(
		"github.com/highlight/highlight/sdk/highlight-go",
		trace.WithInstrumentationVersion("v0.1.0"),
		trace.WithSchemaURL(semconv.SchemaURL),
	)
)

func StartOTLP() (*OTLP, error) {
	var options []otlptracehttp.Option
	if strings.HasPrefix(conf.otlpEndpoint, "http://") {
		options = append(options, otlptracehttp.WithEndpoint(conf.otlpEndpoint[7:]), otlptracehttp.WithInsecure())
	} else if strings.HasPrefix(conf.otlpEndpoint, "https://") {
		options = append(options, otlptracehttp.WithEndpoint(conf.otlpEndpoint[8:]))
	} else {
		logger.Errorf("an invalid otlp endpoint was configured %s", conf.otlpEndpoint)
	}
	options = append(options, otlptracehttp.WithCompression(otlptracehttp.GzipCompression))
	client := otlptracehttp.NewClient(options...)
	exporter, err := otlptrace.New(context.Background(), client)
	if err != nil {
		return nil, fmt.Errorf("creating OTLP trace exporter: %w", err)
	}
	resources, err := resource.New(context.Background(),
		resource.WithFromEnv(),
		resource.WithHost(),
		resource.WithContainer(),
		resource.WithOS(),
		resource.WithProcess(),
		resource.WithAttributes(conf.resourceAttributes...),
	)
	if err != nil {
		return nil, fmt.Errorf("creating OTLP resource context: %w", err)
	}
	h := &OTLP{
		tracerProvider: sdktrace.NewTracerProvider(
			sdktrace.WithSampler(getSampler()),
			sdktrace.WithBatcher(
				exporter,
				sdktrace.WithBatchTimeout(1000*time.Millisecond),
				sdktrace.WithMaxExportBatchSize(128),
				sdktrace.WithMaxQueueSize(1024)),
			sdktrace.WithResource(resources),
		),
	}
	otel.SetTracerProvider(h.tracerProvider)
	return h, nil
}

func (o *OTLP) shutdown() {
	err := o.tracerProvider.ForceFlush(context.Background())
	if err != nil {
		logger.Error(err)
	}
	err = o.tracerProvider.Shutdown(context.Background())
	if err != nil {
		logger.Error(err)
	}
}

func StartTraceWithTimestamp(ctx context.Context, name string, t time.Time, opts []trace.SpanStartOption, tags ...attribute.KeyValue) (trace.Span, context.Context) {
	sessionID, requestID, _ := validateRequest(ctx)
	spanCtx := trace.SpanContextFromContext(ctx)
	if requestID != "" {
		data, _ := base64.StdEncoding.DecodeString(requestID)
		hex := fmt.Sprintf("%032x", data)
		tid, _ := trace.TraceIDFromHex(hex)
		spanCtx = spanCtx.WithTraceID(tid)
	}
	opts = append(opts, trace.WithTimestamp(t))
	ctx, span := tracer.Start(trace.ContextWithSpanContext(ctx, spanCtx), name, opts...)
	span.SetAttributes(
		attribute.String(ProjectIDAttribute, conf.projectID),
		attribute.String(SessionIDAttribute, sessionID),
		attribute.String(RequestIDAttribute, requestID),
	)
	// prioritize values passed in tags for project, session, request ids
	span.SetAttributes(tags...)
	return span, ctx
}

func StartTrace(ctx context.Context, name string, tags ...attribute.KeyValue) (trace.Span, context.Context) {
	return StartTraceWithTimestamp(ctx, name, time.Now(), nil, tags...)
}

func StartTraceWithoutResourceAttributes(ctx context.Context, name string, opts []trace.SpanStartOption, tags ...attribute.KeyValue) (trace.Span, context.Context) {
	resourceAttributes := []attribute.KeyValue{
		semconv.ServiceNameKey.String(""),
		semconv.ServiceVersionKey.String(""),
		semconv.ContainerIDKey.String(""),
		semconv.HostNameKey.String(""),
		semconv.OSDescriptionKey.String(""),
		semconv.OSTypeKey.String(""),
		semconv.ProcessExecutableNameKey.String(""),
		semconv.ProcessExecutablePathKey.String(""),
		semconv.ProcessOwnerKey.String(""),
		semconv.ProcessPIDKey.String(""),
		semconv.ProcessRuntimeDescriptionKey.String(""),
		semconv.ProcessRuntimeNameKey.String(""),
		semconv.ProcessRuntimeVersionKey.String(""),
	}

	attrs := append(resourceAttributes, tags...)

	return StartTraceWithTimestamp(ctx, name, time.Now(), opts, attrs...)
}

func EndTrace(span trace.Span) {
	span.End(trace.WithStackTrace(true))
}

// RecordMetric is used to record arbitrary metrics in your golang backend.
// Highlight will process these metrics in the context of your session and expose them
// through dashboards. For example, you may want to record the latency of a DB query
// as a metric that you would like to graph and monitor. You'll be able to view the metric
// in the context of the session and network request and recorded it.
func RecordMetric(ctx context.Context, name string, value float64, tags ...attribute.KeyValue) {
	span, _ := StartTraceWithTimestamp(ctx, "highlight-metric", time.Now(), []trace.SpanStartOption{trace.WithSpanKind(trace.SpanKindClient)}, tags...)
	defer EndTrace(span)
	span.AddEvent(MetricEvent, trace.WithAttributes(attribute.String(MetricEventName, name), attribute.Float64(MetricEventValue, value)))
}

// RecordError processes `err` to be recorded as a part of the session or network request.
// Highlight session and trace are inferred from the context.
// If no sessionID is set, then the error is associated with the project without a session context.
func RecordError(ctx context.Context, err error, tags ...attribute.KeyValue) context.Context {
	span, ctx := StartTraceWithTimestamp(ctx, "highlight-ctx", time.Now(), []trace.SpanStartOption{trace.WithSpanKind(trace.SpanKindClient)}, tags...)
	defer EndTrace(span)
	RecordSpanError(span, err)
	return ctx
}

func RecordSpanError(span trace.Span, err error, tags ...attribute.KeyValue) {
	if urlErr, ok := err.(*url.Error); ok {
		span.SetAttributes(attribute.String("Op", urlErr.Op))
		span.SetAttributes(attribute.String(ErrorURLAttribute, urlErr.URL))
	}
	span.SetAttributes(tags...)
	// if this is an error with true stacktrace, then create the event directly since otel doesn't support saving a custom stacktrace
	var stackErr ErrorWithStack
	if errors.As(err, &stackErr) {
		RecordSpanErrorWithStack(span, stackErr)
	} else {
		span.RecordError(err, trace.WithStackTrace(true))
	}
}

func RecordSpanErrorWithStack(span trace.Span, err ErrorWithStack) {
	stackTrace := fmt.Sprintf("%+v", err.StackTrace())
	span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(
		semconv.ExceptionTypeKey.String(reflect.TypeOf(err).String()),
		semconv.ExceptionMessageKey.String(err.Error()),
		semconv.ExceptionStacktraceKey.String(stackTrace),
	))
}
