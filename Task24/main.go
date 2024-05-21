package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func distance(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}
func main() {
	point1 := NewPoint(7.2, 9.9)
	point2 := NewPoint(3.7, 1.2)

	fmt.Println(distance(point1, point2))
}
