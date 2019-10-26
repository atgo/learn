package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	ctx, can := context.WithCancel(context.Background())

	server := server(ctx, router())

	go func() {
		err := server.ListenAndServe()
		can()
		panic(err)
	}()

	forever := make(chan os.Signal)
	signal.Notify(forever)
	<-forever // terminated
	can()
}

func server(ctx context.Context, router *mux.Router) http.Server {
	return http.Server{
		Addr:              "localhost:8484",
		Handler:           router,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil, // custom logger
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
		ConnContext: nil, // factory to create connect context.
	}
}

func router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			r.Header.Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"status":"OK"}`))
		},
	)

	return router
}
