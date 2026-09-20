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

func GetAngle(p1, p2 Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y

	// Atan2(dx, -dy) converts Cartesian/screen coordinates
	// into a clockwise compass bearing from north.
	angle := math.Atan2(dx, dy) * 180 / math.Pi

	// Normalize [-180, 180] to [0, 360).
	return math.Mod(angle+360, 360)
}
