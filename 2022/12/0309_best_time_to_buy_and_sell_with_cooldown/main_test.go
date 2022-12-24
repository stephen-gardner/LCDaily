package main

import "testing"

func test(t *testing.T, nums []int, expected int) {
	if res := maxProfit(nums); res != expected {
		t.Log("Input:", nums)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestMaxProfit(t *testing.T) {
	test(t, []int{1, 2, 3, 0, 2}, 3)
	test(t, []int{1}, 0)
}
