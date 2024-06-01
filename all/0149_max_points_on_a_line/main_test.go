package main

import "testing"

func test(t *testing.T, points [][]int, expected int) {
	if res := maxPoints(points); res != expected {
		t.Log("Input:", points)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestMaxPoints(t *testing.T) {
	test(t, [][]int{{1, 1}, {2, 2}, {3, 3}}, 3)
	test(t, [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}, 4)
	test(t, [][]int{{2, 3}, {3, 3}, {-5, 3}}, 3)
	test(t, [][]int{{0, 1}, {0, 0}, {0, 4}, {0, -2}, {0, -1}, {0, 3}, {0, -4}}, 7)
}
