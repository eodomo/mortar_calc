package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("POST /", calculateHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
