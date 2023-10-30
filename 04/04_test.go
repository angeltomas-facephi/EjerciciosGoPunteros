package main

import (
	"testing"
)

func TestFindMajor(t *testing.T) {

	var useCases = []struct {
		name  string
		input []int
		major int
	}{
		{"One", []int{2, 10, 11, 6, 4}, 2},
		{"Two", []int{6, 5, 2, 1, 4}, 0},
		{"Three", []int{9, 8, 7, 10, 6}, 3},
		{"Four", []int{110, 101, 130, 120, 140}, 4},
	}

	for i, element := range useCases {
		memory := findMajor(element.input)

		if memory == &element.input[0] {
			t.Logf(`Test %s passed OK`, element.name)
		} else {
			t.Errorf("Expected: %v, got: %v", &element.input[i], memory)
		}
	}
}
