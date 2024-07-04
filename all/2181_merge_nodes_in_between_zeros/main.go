// Problem: https://leetcode.com/problems/merge-nodes-in-between-zeros/
// Results: https://leetcode.com/problems/merge-nodes-in-between-zeros/submissions/1309712346
package main

import "github.com/stephen-gardner/LCDaily/test/list"

type ListNode = list.ListNode

// Time: O(n)
// Space: O(1)
func mergeNodes(head *ListNode) *ListNode {
	curr := &head
	for tail := head; tail != nil; tail = tail.Next {
		if tail.Val == 0 && (*curr).Val != 0 {
			(*curr).Next = tail
			curr = &(*curr).Next
		} else {
			(*curr).Val += tail.Val
		}
	}
	*curr = nil
	return head
}
