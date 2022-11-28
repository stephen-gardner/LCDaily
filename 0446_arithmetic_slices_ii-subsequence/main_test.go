package main

import "testing"

func test(t *testing.T, nums []int, expected int) {
	if res := numberOfArithmeticSlices(nums); res != expected {
		t.Log("Input:", nums)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestNumberOfArithmeticSlices(t *testing.T) {
	test(t, []int{2, 4, 6, 8, 10}, 7)
	test(t, []int{7, 7, 7, 7, 7}, 16)
}
