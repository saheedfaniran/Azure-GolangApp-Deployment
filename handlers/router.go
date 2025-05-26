package handlers

import (
	"log"
	"time"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	start := time.Now()
	log.Printf("Starting application. Please be patient, this may take a while ...")
	time.Sleep(35 * time.Second)

	r := mux.NewRouter()
	r.HandleFunc("/", helloWorld).Methods("GET")
	r.HandleFunc("/healthz", healthz)
	r.HandleFunc("/greetings", greetings)
	log.Println("Application started in", time.Since(start))
	return r
}
