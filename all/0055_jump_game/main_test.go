package main

import "testing"

func test(t *testing.T, nums []int, expected bool) {
	if canJump(nums) != expected {
		t.Log("Input:", nums)
		t.Fatalf("%v != %v (expected)", !expected, expected)
	}
}

func TestCanJump(t *testing.T) {
	test(t, []int{2, 3, 1, 1, 4}, true)
	test(t, []int{3, 2, 1, 0, 4}, false)
}
