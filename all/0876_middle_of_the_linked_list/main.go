// Problem: https://leetcode.com/problems/middle-of-the-linked-list/
// Results: https://leetcode.com/submissions/detail/854762055/
package main

import "github.com/stephen-gardner/LCDaily/test/list"

type ListNode = list.ListNode

func middleNode(head *ListNode) *ListNode {
	mid := head
	end := head.Next
	for end != nil {
		mid = mid.Next
		end = end.Next
		if end != nil {
			end = end.Next
		}
	}
	return mid
}
