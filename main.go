package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, t)
}

func main() {
	p1 := Point{3, 71}
	p2 := Point{6, 26}

	fmt.Println("Distance: ", GetDistance(p1, p2))
	fmt.Println("Angle: ", GetAngle(p1, p2))

	http.HandleFunc("/index.html", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
