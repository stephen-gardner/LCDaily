package main

import "testing"

func test(t *testing.T, nums, expected []int) {
	dupe := make([]int, len(nums))
	for i := range dupe {
		dupe[i] = nums[i]
	}
	res := removeDuplicates(dupe)
	fail := res != len(expected)
	if !fail {
		for i := range expected {
			if dupe[i] != expected[i] {
				fail = true
				break
			}
		}
	}
	if fail {
		t.Log("   Input:", nums)
		t.Log("  Result:", dupe[:len(expected)])
		t.Log("       K:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestRemoveDuplicates(t *testing.T) {
	test(t, []int{1, 1, 2}, []int{1, 2})
	test(t, []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4})
}
