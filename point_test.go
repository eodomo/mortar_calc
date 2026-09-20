package main

import (
	"testing"
)

func TestGetDistance(t *testing.T) {
	p1 := Point{64, 0}
	p2 := Point{74, 0}

	ans := GetDistance(p1, p2)
	if ans != 10 {
		t.Errorf("GetDistance([64, 0], [75,0]) = %f; want 10", ans)
	}
}

func TestGetAngle(t *testing.T) {
	p1 := Point{0, 0}
	p2 := Point{1, 1}

	ans := GetAngle(p1, p2)
	if ans != 45 {
		t.Errorf("GetAngle([0, 0], [1, 1]) = %f; want 45", ans)
	}
}
