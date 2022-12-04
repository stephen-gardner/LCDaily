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

func listElementAt(head *ListNode, n int) *ListNode {
	for i := 0; i < n; i++ {
		head = head.Next
	}
	return head
}

func test(t *testing.T, nums, expected []int, n int) {
	list := listCreate(nums)
	deleteNode(listElementAt(list, n))
	res := listArray(list)
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

func TestDeleteNode(t *testing.T) {
	test(t, []int{4, 5, 1, 9}, []int{4, 1, 9}, 1)
	test(t, []int{4, 5, 1, 9}, []int{4, 5, 9}, 2)
}
