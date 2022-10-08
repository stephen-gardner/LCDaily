package main

import "testing"

func test(t *testing.T, nums []int, target, expected int) {
	res := threeSumClosest(nums, target)
	if res != expected {
		t.Log(" Input:", nums)
		t.Log("Target:", target)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestThreeSumClosest(t *testing.T) {
	test(t, []int{-1, 2, 1, -4}, 1, 2)
	test(t, []int{0, 0, 0}, 1, 0)
	test(t, []int{3, 9, 11, 12, 18, 42}, 17, 23)
	test(t, []int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2, -2)
	test(t, []int{0, 1, 2}, 0, 3)
	test(t, []int{1, 2, 3}, -21, 6)
}
