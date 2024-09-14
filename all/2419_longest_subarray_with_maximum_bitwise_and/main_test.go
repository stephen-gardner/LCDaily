package main

import "testing"

func TestLongestSubarray(t *testing.T) {
	test := func(nums []int, expected int) {
		if res := longestSubarray(nums); res != expected {
			t.Log("Nums:", nums)
			t.Fatalf("%d != %d (expected)", res, expected)
		}
	}
	test([]int{1, 2, 3, 3, 2, 2}, 2)
	test([]int{1, 2, 3, 4}, 1)
}
