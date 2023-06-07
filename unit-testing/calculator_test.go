package unit_testing

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

//func TestMultiply(t *testing.T) {
//	calculator := Calculator{}
//	result := calculator.Multiply(4, 2)
//	expected := 8
//	if result != expected {
//		t.Errorf("Expected %d, but got %d", expected, result)
//	}
//}

func TestAdd2(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	assert.Equal(t, expected, result)
}
