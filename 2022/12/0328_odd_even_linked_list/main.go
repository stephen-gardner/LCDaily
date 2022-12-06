// Problem: https://leetcode.com/problems/odd-even-linked-list/
// Results: https://leetcode.com/problems/odd-even-linked-list/submissions/855313121/
package main

import "github.com/stephen-gardner/LCDaily/test/list"

type ListNode = list.ListNode

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	oddHead := head.Next
	evenCurr, oddCurr := head, oddHead
	for curr, odd := oddCurr.Next, false; curr != nil; curr = curr.Next {
		if odd {
			oddCurr.Next = curr
			oddCurr = oddCurr.Next
		} else {
			evenCurr.Next = curr
			evenCurr = evenCurr.Next
		}
		odd = !odd
	}
	oddCurr.Next = nil
	evenCurr.Next = oddHead
	return head
}
