package tracer

import "go.opentelemetry.io/otel/attribute"

func SpanTagString(key string, value string) attribute.KeyValue {
	return attribute.Key(key).String(value)
}

func SpanTagInt64(key string, value int64) attribute.KeyValue {
	return attribute.Key(key).Int64(value)
}

func SpanTagBool(key string, value bool) attribute.KeyValue {
	return attribute.Key(key).Bool(value)
}

func SpanTagFloat64(key string, value float64) attribute.KeyValue {
	return attribute.Key(key).Float64(value)
}

func SpanTagStringSlice(key string, value []string) attribute.KeyValue {
	return attribute.Key(key).StringSlice(value)
}

func SpanTagInt(key string, value int) attribute.KeyValue {
	return attribute.Key(key).Int(value)
}
