package main

import "testing"

func test(t *testing.T, nums []int, expected int) {
	if res := minimumAverageDifference(nums); res != expected {
		t.Log("Input:", nums)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestMinimumAverageDifference(t *testing.T) {
	test(t, []int{2, 5, 3, 9, 5, 3}, 3)
	test(t, []int{0}, 0)
	test(t, []int{1}, 0)
}
