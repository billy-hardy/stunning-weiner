package main

import (
	"github.com/gorilla/mux"
	h "handlers"
	"net/http"
)

const (
	GET = "GET"
)

func get(r *mux.Router, path string, f h.Handler) {
	r.HandleFunc(path, h.ErrorHandler(f)).Methods(GET)
}

func main() {

	r := mux.NewRouter()
	get(r, "/", h.RootHandler)
	get(r, "/reverse/{word}", h.HelloHandler)
	listenAndServe(":8080", r)
}

func listenAndServe(addr string, handler http.Handler) error {
	server := &http.Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
