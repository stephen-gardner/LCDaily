package main

import "testing"

func test(t *testing.T, nums []int, expected bool) {
	if increasingTriplet(nums) != expected {
		t.Log("Input:", nums)
		t.Fatalf("%v != %v (expected)", !expected, expected)
	}
}

func TestIncreasingTriplet(t *testing.T) {
	test(t, []int{1, 2, 3, 4, 5}, true)
	test(t, []int{5, 4, 3, 2, 1}, false)
	test(t, []int{2, 1, 5, 0, 4, 6}, true)
}
