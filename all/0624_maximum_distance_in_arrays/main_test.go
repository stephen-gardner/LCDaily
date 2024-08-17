package main

import "testing"

func TestMaxDistance(t *testing.T) {
	test := func(arrays [][]int, expected int) {
		if res := maxDistance(arrays); res != expected {
			t.Log("Input:", arrays)
			t.Fatalf("%d != %d (expected)", res, expected)
		}
	}
	test([][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}, 4)
	test([][]int{{1}, {1}}, 0)
}
