package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Distance(other *Point) float64 {
	d := math.Sqrt(((other.x - p.x) * (other.x - p.x)) + ((other.y - p.y) * (other.y - p.y)))
	return d
}

func main() {
	a := NewPoint(2, 3)
	b := NewPoint(5, 7)
	fmt.Println(a.Distance(b))
}