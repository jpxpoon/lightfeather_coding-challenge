package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = "3001"

var router = &mux.Router{}

// Endpoints:
func registerEndpoints() {
	router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/supervisors", getSupervisors).Methods("GET")
	router.HandleFunc("/api/submit", saveNotification).Methods("POST", "OPTIONS")
}

func main() {
	registerEndpoints()

	log.Printf(`Listening on http://localhost:%s/`, port)

	if err := http.ListenAndServe("0.0.0.0:"+port, router); err != nil {
		panic(fmt.Errorf("start http server: %w", err))
	}
}
