package main

import (
	"testing"

	"github.com/stephen-gardner/LCDaily/test/list"
)

func test(t *testing.T, nums []int) {
	head := list.Create(nums)
	expected := head
	for mid := len(nums) / 2; mid > 0; mid-- {
		expected = expected.Next
	}
	if res := middleNode(head); res != expected {
		t.Log("   Input:", nums)
		if res != nil {
			t.Logf("  Result: %p (%d)", res, res.Val)
		} else {
			t.Log("  Result: nil")
		}
		t.Logf("Expected: %p (%d)", expected, expected.Val)
		t.FailNow()
	}
}

func TestMiddleNode(t *testing.T) {
	test(t, []int{1, 2, 3, 4, 5})
	test(t, []int{1, 2, 3, 4, 5, 6})
	test(t, []int{1})
}
