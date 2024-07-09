// Problem: https://leetcode.com/problems/find-the-winner-of-the-circular-game/
// Results: https://leetcode.com/problems/find-the-winner-of-the-circular-game/submissions/1313438784
package main

type friend struct {
	index int
	next  *friend
}

// Time: O(n*k)
// Space: O(n)
func findTheWinner(n int, k int) int {
	head := &friend{index: 1}
	curr := &head.next
	for i := 2; i <= n; i++ {
		*curr = &friend{index: i}
		curr = &(*curr).next
	}
	*curr = head
	for (*curr).next != *curr {
		for i := 1; i < k; i++ {
			curr = &(*curr).next
		}
		*curr = (*curr).next
	}
	return (*curr).index
}
