package main

import "testing"

func test(t *testing.T, n int, dislikes [][]int, expected bool) {
	if possibleBipartition(n, dislikes) != expected {
		t.Logf("Input: %v; n = %d", dislikes, n)
		t.Fatalf("%v != %v", !expected, expected)
	}
}

func TestPossibleBipartition(t *testing.T) {
	test(t, 4, [][]int{{1, 2}, {1, 3}, {2, 4}}, true)
	test(t, 3, [][]int{{1, 2}, {1, 3}, {2, 3}}, false)
	test(t, 5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}, false)
}
