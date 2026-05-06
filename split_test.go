package main

import "testing"

func TestSplitNoNewline(t *testing.T) {
	result := Split("Hello")
	if len(result) != 1 {
		t.Errorf("Expected 1 part, got %d", len(result))
	}
}

func TestSplitSingleNewline(t *testing.T) {
	result := Split(`Hello\nThere`)
	if len(result) != 2 {
		t.Errorf("Expected 2 parts, got %d", len(result))
	}
}

func TestSplitDoubleNewline(t *testing.T) {
	result := Split(`Hello\n\nThere`)
	if len(result) != 3 {
		t.Errorf("Expected 3 parts, got %d", len(result))
	}
}

func TestSplitEmpty(t *testing.T) {
	result := Split("")
	if len(result) != 1 {
		t.Errorf("Expected 1 part, got %d", len(result))
	}
}
