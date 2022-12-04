// Problem: https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// Results: https://leetcode.com/submissions/detail/810412961/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}
	if length == n {
		if length == 1 {
			return nil
		}
		return head.Next
	}
	curr = head
	for length--; length > n; length-- {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	return head
}
