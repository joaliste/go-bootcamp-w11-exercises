package tools_test

import (
	"test-code/tools"
	"testing"
)

func TestElementAtIndex_ExistentIndex(t *testing.T) {
	// Arrange
	slice := []int{1, 2, 3, 4, 5}
	index := 0

	expected := 1
	// Act
	result, err := tools.ElementAtIndex(slice, index)

	// Assert
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
	if result != expected {
		t.Fatalf("Expected %d, got %d", expected, result)
	}
	t.Log("Success")
}
func TestElementAtIndex_NonExistingIndex(t *testing.T) {
	// Arrange
	slice := []int{1, 2, 3, 4, 5}
	idx := len(slice)

	expected := 5
	// Act
	result, err := tools.ElementAtIndex(slice, idx)
	// Assert
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}

	if result != expected {
		t.Fatalf("Test Failed. Expected error but %d was given", result)
	}
}
