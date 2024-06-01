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

func test(t *testing.T, n int, nums, expected []int) {
	res := listArray(removeNthFromEnd(listCreate(nums), n))
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

func TestRemoveNthFromEnd(t *testing.T) {
	test(t, 2, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 5})
	test(t, 1, []int{1}, []int{})
	test(t, 1, []int{1, 2}, []int{1})
}
