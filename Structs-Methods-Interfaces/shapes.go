package main

// Perimeter calculation
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

// Area calculator
func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}
