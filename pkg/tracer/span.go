package tracer

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type Span struct {
	span trace.Span
}

// StartSpan returns a new span from the global tracer. Depending on the `cus`
// argument, the span is either a plain one or a customised one. Each resulting
// span must be completed with `defer span.End()` right after the call.
func SpanStart(ctx context.Context, name string, tags ...attribute.KeyValue) (context.Context, *Span) {
	ctx, span := otel.Tracer("").Start(ctx, name, trace.WithAttributes(tags...))

	return ctx, &Span{span: span}
}

func (s *Span) Finish() {
	s.span.End()
}

// AddSpanTags adds a new tags to the span. It will appear under "Tags" section
// of the selected span. Use this if you think the tag and its value could be
// useful while debugging.
func (s *Span) AddTags(tags ...attribute.KeyValue) {
	s.span.SetAttributes(tags...)
}

// AddSpanEvents adds a new events to the span. It will appear under the "Logs"
// section of the selected span. Use this if the event could mean anything
// valuable while debugging.
func (s *Span) AddEvents(name string, attrs ...attribute.KeyValue) {
	s.span.AddEvent(name, trace.WithAttributes(attrs...))
}

// AddSpanError adds a new event to the span. It will appear under the "Logs"
// section of the selected span. This is not going to flag the span as "failed".
// Use this if you think you should log any exceptions such as critical, error,
// warning, caution etc. Avoid logging sensitive data!
func (s *Span) AddError(err error, attrs ...attribute.KeyValue) {
	s.span.RecordError(err, trace.WithAttributes(attrs...))
}

func (s *Span) TraceId() string {
	return s.span.SpanContext().TraceID().String()
}

func (s *Span) AddAnyTags(keyValues map[string]any) {
	fields := make([]attribute.KeyValue, len(keyValues))

	for key, val := range keyValues {
		switch typedVal := val.(type) {
		case bool:
			fields = append(fields, attribute.Bool(key, typedVal))
		case string:
			fields = append(fields, attribute.String(key, typedVal))
		case int:
			fields = append(fields, attribute.Int(key, typedVal))
		case int8:
			fields = append(fields, attribute.Int(key, int(typedVal)))
		case int16:
			fields = append(fields, attribute.Int(key, int(typedVal)))
		case int32:
			fields = append(fields, attribute.Int(key, int(typedVal)))
		case int64:
			fields = append(fields, attribute.Int64(key, typedVal))
		case uint:
			fields = append(fields, attribute.Int(key, int(typedVal)))
		case uint64:
			fields = append(fields, attribute.Int(key, int(typedVal)))
		case uint8:
			fields = append(fields, attribute.Int(key, int(typedVal)))
		case uint16:
			fields = append(fields, attribute.Int(key, int(typedVal)))
		case uint32:
			fields = append(fields, attribute.Int(key, int(typedVal)))
		case float32:
			fields = append(fields, attribute.Float64(key, float64(typedVal)))
		case float64:
			fields = append(fields, attribute.Float64(key, typedVal))
		default:
			// When in doubt, coerce to a string
			fields = append(fields, attribute.String(key, fmt.Sprintf("%v", typedVal)))
		}
	}

	s.span.SetAttributes(fields...)
}
