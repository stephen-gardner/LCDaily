package main

import "testing"

func test(t *testing.T, stones [][]int, expected int) {
	if res := removeStones(stones); res != expected {
		t.Log("   Input:", stones)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestRemoveStones(t *testing.T) {
	test(t, [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}, 5)
	test(t, [][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}, 3)
	test(t, [][]int{{0, 0}}, 0)
}
