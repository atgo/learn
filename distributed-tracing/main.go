package distributed_tracing

import (
	"io/ioutil"
	"log"

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
