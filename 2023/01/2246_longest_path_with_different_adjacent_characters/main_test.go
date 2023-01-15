package main

import "testing"

func test(t *testing.T, parent []int, s string, expected int) {
	if res := longestPath(parent, s); res != expected {
		t.Log("     s:", s)
		t.Log("Parent:", parent)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestLongestPath(t *testing.T) {
	test(t, []int{-1, 0, 0, 1, 1, 2}, "abacbe", 3)
	test(t, []int{-1, 0, 0, 0}, "aabc", 3)
	test(t, []int{-1, 0, 1}, "aab", 2)
}
