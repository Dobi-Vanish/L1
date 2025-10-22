package main

import (
	"fmt"
	"math"
)

// Point представляет точку на плоскости с приватными полями
type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point1 := NewPoint(1.0, 3.0)
	point2 := NewPoint(4.0, 6.0)

	distance1 := point1.Distance(point2)

	fmt.Printf("Point 1: %f\n", point1)
	fmt.Printf("Point 2: %f\n", point2)
	fmt.Printf("Distance between: %.2f\n\n", distance1)

}
