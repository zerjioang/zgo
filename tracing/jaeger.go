//
// Copyright zerjioang. 2021 All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package tracing

import (
	"context"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"

	"github.com/uber/jaeger-lib/metrics"

	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

// docs: https://logz.io/blog/go-instrumentation-distributed-tracing-jaeger/
var (
	tracer opentracing.Tracer
	closer io.Closer
)

// Start opens the tracer
func Start(name string) error {
	// Sample configuration for testing. Use constant sampling to sample every trace
	// and enable LogSpan to log every span via configured Logger.
	cfg := jaegercfg.Configuration{
		ServiceName: name,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}
	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	var err error
	tracer, closer, err = cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		return err
	}
	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)
	return nil
}

// Close closes the tracer
func Close() error {
	if closer != nil {
		return closer.Close()
	}
	return nil
}

// StartSpan
// Span creates and starts a new trace using global setup tracer
// receives span name and returns it
// Note: This method assumes jaeger.Start was called previously
func StartSpan(name string) opentracing.Span {
	gtracer := opentracing.GlobalTracer()
	span := gtracer.StartSpan(name)
	return span
}

// StartChildSpan creates and starts a new trace using global setup tracer
// The created span is a child of passed Span
// receives span name and returns it
// Note: This method assumes jaeger.Start was called previously
func StartChildSpan(name string, parentCtx context.Context) opentracing.Span {
	if parentCtx == nil {
		// it might be a developer error, that forget to define the parent context
		// so we return a span with no parent
		return StartSpan(name)
	}
	// extract parent span from context
	parentVal := parentCtx.Value("span")
	if parentVal == nil {
		return StartSpan(name)
	}
	parent, ok := parentVal.(opentracing.Span)
	if !ok {
		return StartSpan(name)
	}
	gtracer := opentracing.GlobalTracer()
	childSpan := gtracer.StartSpan(
		name,
		opentracing.ChildOf(parent.Context()),
	)
	return childSpan
}
