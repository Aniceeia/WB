package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) Distance(other Point) float64 {
	newX := p.x - other.x
	newY := p.y - other.y
	return math.Sqrt(newX*newX + newY*newY)
}

func main() {
	a1 := NewPoint(1.0, 2.0)
	a2 := NewPoint(3.0, 4.0)

	distance := a1.Distance(a2)

	fmt.Printf("distance: %.f", distance)
}
