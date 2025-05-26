package main

import (
	"flag"
	"log"
	"net/http"

	"techtest-devops/handlers"
)

func main() {
	addrFlag := flag.String("addr", ":1337", "server address:port")
	flag.Parse()
	router := handlers.Router()
	log.Fatal(http.ListenAndServe(*addrFlag, router))
}
