package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"stringutils"
	"time"
)

type Handler func(http.ResponseWriter, *http.Request) error

func RootHandler(w http.ResponseWriter, r *http.Request) error {
	return ServeContent(w, r, "<h1>Server is running.</h1>", http.StatusOK)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	return ServeContent(w, r, stringutils.Reverse(params["word"]), http.StatusOK)
}

func ErrorHandler(f Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			ServeContent(w, r, err.Error(), http.StatusInternalServerError)
		}
	}
}

func ServeContent(w http.ResponseWriter, r *http.Request, body string, statusCode int) error {
	logHeader := fmt.Sprintf(time.Now().Format("Jan 2 15:04:05 EST"))
	path := fmt.Sprintf(", request on path '%s' ", r.URL.Path)
	status := fmt.Sprintf("returned status code: %d", statusCode)

	if statusCode > 299 {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		errString := fmt.Sprintf(" with error: %s", body)
		status = status + errString
	}

	log.Println(logHeader + path + status)
	w.WriteHeader(statusCode)
	w.Write([]byte(body))

	return nil
}
