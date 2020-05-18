package main

import "math"

// A Circle type
type Circle struct {
	radius float64
}

// Area Calculation function
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
