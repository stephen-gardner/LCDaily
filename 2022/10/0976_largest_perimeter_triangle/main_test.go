package main

import "testing"

func test(t *testing.T, nums []int, expected int) {
	if res := largestPerimeter(nums); res != expected {
		t.Log("Input:", nums)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestLargestPerimeter(t *testing.T) {
	test(t, []int{2, 1, 2}, 5)
	test(t, []int{1, 2, 1}, 0)
	test(t, []int{3, 9, 11, 12, 18, 42}, 41)
}
