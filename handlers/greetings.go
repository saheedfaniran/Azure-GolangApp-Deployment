package handlers

import (
	"fmt"
	"net/http"
	"os"
)

// Get preferred outbound ip of this machine
func greetings(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("CANDIDATE_FIRSTNAME")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello %s, welcome to Hello World!\n", name)
}
