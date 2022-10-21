package main

import "testing"

func test(t *testing.T, nums []int, k int, expected bool) {
	if containsNearbyDuplicate(nums, k) != expected {
		t.Log("k =", k)
		t.Log("Input:", nums)
		t.Fatalf("%v != %v (expected)", !expected, expected)
	}
}

func TestContainsNearbyDuplicate(t *testing.T) {
	test(t, []int{1, 2, 3, 1}, 3, true)
	test(t, []int{1, 0, 1, 1}, 1, true)
	test(t, []int{1, 2, 3, 1, 2, 3}, 2, false)
	test(t, []int{99, 99}, 2, true)
}
