package main

import "testing"

func listArray(curr *ListNode) []int {
	arr := []int{}
	for curr != nil {
		arr = append(arr, curr.Val)
		curr = curr.Next
	}
	return arr
}

func listCreate(nums []int) *ListNode {
	var head *ListNode
	curr := &head
	for _, n := range nums {
		*curr = &ListNode{Val: n}
		curr = &(*curr).Next
	}
	return head
}

func test(t *testing.T, nums, expected []int) {
	list := listCreate(nums)
	res := listArray(deleteMiddle(list))
	fail := len(res) != len(expected)
	if !fail {
		for i := range expected {
			if res[i] != expected[i] {
				fail = true
				break
			}
		}
	}
	if fail {
		t.Log("   Input:", nums)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestDeleteMiddle(t *testing.T) {
	test(t, []int{1, 3, 4, 7, 1, 2, 6}, []int{1, 3, 4, 1, 2, 6})
	test(t, []int{1, 2, 3, 4}, []int{1, 2, 4})
	test(t, []int{1}, []int{})
	test(t, []int{1, 2}, []int{1})
}
