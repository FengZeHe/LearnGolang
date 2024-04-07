package test

import "testing"

func TestAdd(t *testing.T) {
	result := Add(3, 5)
	expected := 8

	if result != expected {
		t.Errorf("Add(3,5) returned %d, expected %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("Subtract(5,3) returned %d, expected %d", result, expected)
	}
}
