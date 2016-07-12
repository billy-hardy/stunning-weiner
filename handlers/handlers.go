package handlers

import (
	"fmt"
	"github.com/billy-hardy/ic-weiner/stringutils"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Handler func(http.ResponseWriter, *http.Request) error

func RootHandler(w http.ResponseWriter, r *http.Request) error {
	return ServeContent(w, r, "<h1>Server is running.</h1>", http.StatusOK)
}

func ReverseStringHandler(w http.ResponseWriter, r *http.Request) error {
	r.ParseForm()
	fmt.Println(r.Form["url_long"])
	params := mux.Vars(r)
	return ServeContent(w, r, stringutils.Reverse(params["word"]), http.StatusOK)
}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) error {
	page, err := ioutil.ReadFile("html/login.html")
	if err != nil {
		return ServeContent(w, r, "page not found", http.StatusNotFound)
	}
	return ServeContent(w, r, string(page[:]), http.StatusOK)
}

func RequestLoginHandler(w http.ResponseWriter, r *http.Request) error {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if len(username) == 0 {
		return ServeContent(w, r, "Username is required, please try login again", http.StatusBadRequest)
	}
	if len(password) == 0 {
		return ServeContent(w, r, "Password is required, please try login again", http.StatusBadRequest)
	}
	log.Printf("user %v logged in with password %v", username, password)
	return ServeContent(w, r, "Login Successful", http.StatusOK)
}

func MoveHandler(w http.ResponseWriter, r *http.Request) error {
	return ServeContent(w, r, "", http.StatusOK)
}

func GameStatusHandler(w http.ResponseWriter, r *http.Request) error {
	return ServeContent(w, r, "", http.StatusOK)
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
