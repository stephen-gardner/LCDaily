package main

import "testing"

func test(t *testing.T, nums []int, k int, expected bool) {
	if checkSubarraySum(nums, k) != expected {
		t.Log("Input:", nums)
		t.Fatalf("%v != %v (expected)", !expected, expected)
	}
}

func TestCheckSubarraySum(t *testing.T) {
	test(t, []int{23, 2, 4, 6, 7}, 6, true)
	test(t, []int{23, 2, 6, 4, 7}, 6, true)
	test(t, []int{23, 2, 6, 4, 7}, 13, false)
	test(t, []int{9, 23, 2, 6, 4, 7}, 6, true)
	test(t, []int{5, 0, 0, 0}, 3, true)
	test(t, []int{0}, 1, false)
	test(t, []int{1, 0}, 2, false)

}
