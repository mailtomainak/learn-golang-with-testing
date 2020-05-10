package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 10, Height: 20}
	actualPerimeter := Perimeter(rectangle)
	expectedPerimeter := 60.00
	if actualPerimeter != expectedPerimeter {
		t.Errorf("Expected %.2f but got %.2f", expectedPerimeter, actualPerimeter)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{Width: 10, Height: 20}
	actualArea := Area(rectangle)
	expectedArea := 200.00

	if actualArea != expectedArea {
		t.Errorf("Expected %.2f but got %.2f", expectedArea, actualArea)
	}

}
