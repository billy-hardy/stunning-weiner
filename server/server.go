package main

import (
	h "github.com/billy-hardy/ic-weiner/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	GET  = "GET"
	POST = "POST"
)

func get(r *mux.Router, path string, f h.Handler) {
	r.HandleFunc(path, h.ErrorHandler(f)).Methods(GET)
}

func post(r *mux.Router, path string, f h.Handler) {
	r.HandleFunc(path, h.ErrorHandler(f)).Methods(POST)
}

func main() {
	r := mux.NewRouter()
	get(r, "/", h.RootHandler)
	get(r, "/login", h.LoginPageHandler)
	post(r, "/login", h.RequestLoginHandler)
	get(r, "/reverse/{word}", h.ReverseStringHandler)
	get(r, "/error}", h.ErrorTestHandler)
	log.Fatal(listenAndServe(":8080", r))
}

func listenAndServe(addr string, handler http.Handler) error {
	server := &http.Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
