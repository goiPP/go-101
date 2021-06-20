package main

import "testing"

func TestGetAreaTriangle(t *testing.T) {
	triangle := triangle{2, 3}

	if triangle.getArea() != 3 {
		t.Errorf("Expected triangle area to be 3, but got %v", triangle.getArea())
	}
}

func TestGetAreaSquare(t *testing.T) {
	square := square{3}

	if square.getArea() != 9 {
		t.Errorf("Expected square area to be 9, but got %v", square.getArea())
	}
}
