// Problem: https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/
// Results: https://leetcode.com/submissions/detail/822007803/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	if head.Next.Next == nil {
		head.Next = nil
		return head
	}
	slow, fast := head, head.Next
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	*slow = *slow.Next
	return head
}
