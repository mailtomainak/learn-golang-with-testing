package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 10, Height: 20}
	actualPerimeter := Perimeter(rectangle)
	expectedPerimeter := 60.00
	if actualPerimeter != expectedPerimeter {
		t.Errorf("Expected %.2f but got %.2f", expectedPerimeter, actualPerimeter)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		want := 72.0
		if got != want {
			t.Errorf("got %g expected %g", got, want)
		}

	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %g expected %g", got, want)
		}
	})

}
