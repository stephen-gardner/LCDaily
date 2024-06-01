package main

import "testing"

func test(t *testing.T, points [][]int, expected int) {
	dupe := make([][]int, len(points))
	for i := range dupe {
		dupe[i] = make([]int, len(points[i]))
		copy(dupe[i], points[i])
	}
	if res := findMinArrowShots(dupe); res != expected {
		t.Log("Input:", points)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestFindMinArrowShots(t *testing.T) {
	test(t, [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2)
	test(t, [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4)
	test(t, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2)
	test(t, [][]int{{9, 12}, {1, 10}, {4, 11}, {8, 12}, {3, 9}, {6, 9}, {6, 7}}, 2)
}
