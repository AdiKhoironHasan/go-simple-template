package tracer

import (
	"context"

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
