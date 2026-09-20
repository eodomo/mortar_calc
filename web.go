package main

import (
	"html/template"
	"net/http"
	"strconv"
)

type OutputData struct {
	P1x      float64
	P1y      float64
	P2x      float64
	P2y      float64
	Distance float64
	Angle    float64
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	data := OutputData{
		P1x:      64,
		P1y:      64,
		P2x:      64,
		P2y:      64,
		Distance: 0,
		Angle:    0,
	}
	t.Execute(w, data)
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "could not parse form", http.StatusBadRequest)
		return
	}

	p1xText := r.Form.Get("p1x")
	p1x, err := strconv.ParseFloat(p1xText, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p1yText := r.Form.Get("p1y")
	p1y, err := strconv.ParseFloat(p1yText, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p1 := Point{p1x, p1y}

	p2xText := r.Form.Get("p2x")
	p2x, err := strconv.ParseFloat(p2xText, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p2yText := r.Form.Get("p2y")
	p2y, err := strconv.ParseFloat(p2yText, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p2 := Point{p2x, p2y}

	Distance := GetDistance(p1, p2)
	Angle := GetAngle(p1, p2)
	t, _ := template.ParseFiles("index.html")

	data := OutputData{
		P1x:      p1.x,
		P1y:      p1.y,
		P2x:      p2.x,
		P2y:      p2.y,
		Distance: Distance,
		Angle:    Angle,
	}
	t.Execute(w, data)
}
