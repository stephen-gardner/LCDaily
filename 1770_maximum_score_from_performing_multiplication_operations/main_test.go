package main

import "testing"

func test(t *testing.T, nums, multipliers []int, expected int) {
	res := maximumScore(nums, multipliers)
	if res != expected {
		t.Log("Nums:", nums)
		t.Log("Multipliers:", multipliers)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestMaximumScore(t *testing.T) {
	test(t, []int{1, 2, 3}, []int{3, 2, 1}, 14)
	test(t, []int{-5, -3, -3, -2, 7, 1}, []int{-10, -5, 3, 4, 6}, 102)
}
