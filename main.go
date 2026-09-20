package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	http.HandleFunc("/", handler)
	http.HandleFunc("POST /", calculateHandler)

	fmt.Println("Starting service on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
