package main

import "testing"

func test(t *testing.T, matrix [][]int, expected bool) {
	if isToeplitzMatrix(matrix) != expected {
		for _, row := range matrix {
			t.Log(row)
		}
		t.Fatalf("%v != %v (expected)", !expected, expected)
	}
}

func TestIsToeplitzMatrix(t *testing.T) {
	test(t, [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}, true)
	test(t, [][]int{{1, 2}, {2, 2}}, false)
}
