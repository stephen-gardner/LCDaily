package main

import (
	"testing"

	"github.com/stephen-gardner/LCDaily/test/list"
)

func test(t *testing.T, nums, expected []int) {
	lst := list.Create(nums)
	res := oddEvenList(lst)
	arr := list.Array(res)
	fail := len(arr) != len(expected)
	if !fail {
		for i := range expected {
			if arr[i] != expected[i] {
				fail = true
				break
			}
		}
	}
	if fail {
		t.Log("   Input:", nums)
		t.Log("  Result:", arr)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestOddEvenList(t *testing.T) {
	test(t, []int{1, 2, 3, 4, 5}, []int{1, 3, 5, 2, 4})
	test(t, []int{2, 1, 3, 5, 6, 4, 7}, []int{2, 3, 6, 7, 1, 5, 4})
	test(t, []int{1, 2}, []int{1, 2})
	test(t, []int{0}, []int{0})
}
