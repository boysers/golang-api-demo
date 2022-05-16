package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type user struct {
	Name  string `json:"full_name"`
	Email string `json:"email_address"`
}

type apiHandler struct{}

func homeHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Hello from Handler %s", req.URL)
}

// http.Handler
func (apiHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	u := user{Name: "boysers le technicien", Email: "boysers@example.com"}
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(u)
}

func withlogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(rw, req)
		end := time.Since(start)
		fmt.Printf("%s %s processing time %s\n", req.Method, req.URL, end)
	})
}

func main() {
	mux := http.DefaultServeMux

	mux.Handle("/home", withlogger(http.HandlerFunc(homeHandler)))
	mux.Handle("/api", withlogger(apiHandler{}))

	http.ListenAndServe(":3000", mux)
}
