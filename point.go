package main

import "math"

type Point struct {
	x float64
	y float64
}

func GetDistance(p1 Point, p2 Point) float64 {
	x2 := math.Pow(p2.x-p1.x, 2)
	y2 := math.Pow(p2.y-p1.y, 2)
	return math.Pow(x2+y2, (0.5)) * 100.0
}

func GetAngle(p1 Point, p2 Point) float64 {
	angle := math.Atan((p2.y-p1.y)/(p2.x-p1.x)) * (180 / math.Pi)
	return math.Mod(math.Mod(angle, 360.0)+360.0, 360.0)
}
