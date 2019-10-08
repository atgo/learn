package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	// Start the tracer and defer the Stop method.
	tracer.Start(tracer.WithAgentAddr("host:port"))
	defer tracer.Stop()

	// Start a root span.
	span := tracer.StartSpan("get.data")
	defer span.Finish()

	// Create a child of it, computing the time needed to read a file.
	child := tracer.StartSpan("read.file", tracer.ChildOf(span.Context()))
	child.SetTag(ext.ResourceName, "test.json")

	// Perform an operation.
	_, err := ioutil.ReadFile("~/test.json")

	// We may finish the child span using the returned error. If it's
	// nil, it will be disregarded.
	child.Finish(tracer.WithError(err))
	if err != nil {
		log.Fatal(err)
	}
}

func withHttpRequest(ctx context.Context) {
	span, ctx := tracer.StartSpanFromContext(ctx, "service.api")
	defer span.Finish()

	go func() {
		parentSpan, _ := tracer.SpanFromContext(ctx)
		span = tracer.StartSpan("service.api.request", tracer.ChildOf(parentSpan.Context()))
		ctx = tracer.ContextWithSpan(ctx, span)

		req, _ := http.NewRequestWithContext(ctx, "POST", "http://x/", nil)
		req.Header.Set("x-datadog-trace-id", string(span.Context().TraceID()))
		req.Header.Set("x-datadog-parent-id", string(span.Context().SpanID()))
		client := http.Client{}
		res, err := client.Do(req)
		if nil != err {
			span.Finish(tracer.WithError(err))
		} else {
			span.SetTag("service", "x")
			span.SetTag("httpStatus", res.StatusCode)
			span.Finish()
		}
	}()
}
