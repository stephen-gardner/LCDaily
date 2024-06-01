package main

import "testing"

func test(t *testing.T, nums []int, expected int) {
	if res := rob(nums); res != expected {
		t.Log("Input:", nums)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestRob(t *testing.T) {
	test(t, []int{1, 2, 3, 1}, 4)
	test(t, []int{2, 7, 9, 3, 1}, 12)
}
