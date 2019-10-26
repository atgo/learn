HTTP server
====

```
go get github.com/gorilla/mux@latest
```

### Example

```golang
package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	server := http.Server{
		Addr:              "localhost:8484",
		Handler:           router(),
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil, // custom logger
		BaseContext:       nil,
		ConnContext:       nil, // factory to create connect context.
	}

	panic(server.ListenAndServe())
}

func router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"status":"OK""}`))
		},
	)

	return router
}
```
