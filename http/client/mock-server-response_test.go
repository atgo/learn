package consumer

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Context(t *testing.T) {
	payload := Payload{
		RoutingKey: "foo",
		Body:       "foo",
		Context: map[string]interface{}{
			"x-datadog-trace-id":          "xxxxx-1",
			"x-datadog-parent-id":         "xxxxx-2",
			"x-datadog-sampling-priority": 333,
			"x-datadog-origin":            "xxxxx-4",
		},
	}

	var log *http.Request

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log = r
		}))
	defer ts.Close()

	_ = payload.push(http.DefaultClient, ts.URL)
	ass := assert.New(t)
	ass.Equal("xxxxx-1", log.Header.Get("X-Datadog-Trace-Id"))
	ass.Equal("xxxxx-2", log.Header.Get("X-Datadog-Parent-Id"))
	ass.Equal("", log.Header.Get("X-Datadog-Sampling-Priority"), "should not have error if failed converting to string")
	ass.Equal("xxxxx-4", log.Header.Get("X-Datadog-Origin"))
}
