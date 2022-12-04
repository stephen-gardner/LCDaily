package main

import "testing"

func test(t *testing.T, nums []int, expected bool) {
	if uniqueOccurrences(nums) != expected {
		t.Log("Input:", nums)
		t.Fatalf("%v != %v (expected)", !expected, expected)
	}
}

func TestUniqueOccurrences(t *testing.T) {
	test(t, []int{1, 2, 2, 1, 1, 3}, true)
	test(t, []int{1, 2}, false)
	test(t, []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true)
}
