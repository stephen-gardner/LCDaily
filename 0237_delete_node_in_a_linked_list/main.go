// Problem: https://leetcode.com/problems/delete-node-in-a-linked-list/
// Results: https://leetcode.com/submissions/detail/821359344/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	*node = *node.Next
}
