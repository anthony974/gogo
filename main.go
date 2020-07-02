package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(WebServer)
	if err := http.ListenAndServe(":80", handler); err != nil {
		log.Fatalf("could not listen on port 80 %v", err)
	}
}
