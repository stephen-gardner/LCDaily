package main

import "testing"

func test(t *testing.T, arr []int, expected int) {
	if res := sumSubarrayMins(arr); res != expected {
		t.Log("Input:", arr)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestSumSubarrayMins(t *testing.T) {
	test(t, []int{3, 1, 2, 4}, 17)
	test(t, []int{11, 81, 94, 43, 3}, 444)
}
