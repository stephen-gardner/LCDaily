// Problem: https://leetcode.com/problems/odd-even-linked-list/
// Results: https://leetcode.com/problems/odd-even-linked-list/submissions/855313121/
package main

import "github.com/stephen-gardner/LCDaily/test/list"

type ListNode = list.ListNode

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	headOdd := head.Next
	currEven, currOdd := head, headOdd
	for curr, odd := currOdd.Next, false; curr != nil; curr = curr.Next {
		if odd {
			currOdd.Next = curr
			currOdd = currOdd.Next
		} else {
			currEven.Next = curr
			currEven = currEven.Next
		}
		odd = !odd
	}
	currEven.Next = headOdd
	currOdd.Next = nil
	return head
}
