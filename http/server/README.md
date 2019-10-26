HTTP server
====

## Mux

- go get github.com/gorilla/mux@latest
- https://github.com/gorilla/mux

## Mux & datadog

- go get gopkg.in/DataDog/dd-trace-go.v1/ddtrace
- https://godoc.org/gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux

```go
mux := muxtrace.NewRouter()
mux.HandleFunc("/", handler)
http.ListenAndServe(":8484", mux)

# with service name
mux := muxtrace.NewRouter(muxtrace.WithServiceName("mux.route"))
mux.HandleFunc("/", handler)
http.ListenAndServe(":8484", mux)
```
