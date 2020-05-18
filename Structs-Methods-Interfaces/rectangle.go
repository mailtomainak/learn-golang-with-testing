package main

// Rectangle type
type Rectangle struct {
	Width  float64
	Height float64
}

// Area function
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
