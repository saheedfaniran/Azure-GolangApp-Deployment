package handlers

import (
	"log"
	"net/http"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("Healthz endpoint returned 200")
}
