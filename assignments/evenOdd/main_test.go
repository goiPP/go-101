package main

import "testing"

func TestGenerateSliceOfSquentialInt(t *testing.T) {
	slice := generateSliceOfSquentialInt(1, 10)

	if len(slice) != 10 {
		t.Errorf("Expected slice legnth of 10, but got %v", len(slice))
	}

	if slice[0] != 1 {
		t.Errorf("Expected first number to be 1, but got %v", slice[0])
	}

	if slice[9] != 10 {
		t.Errorf("Expected last number to be 10, but got %v", slice[0])
	}
}

func TestIsEvent(t *testing.T) {
	if !isEven(8) {
		t.Errorf("Expected that 8 is even, but was not")
	}

	if isEven(1) {
		t.Errorf("Expected that 1 is odd, but was not")
	}
}
