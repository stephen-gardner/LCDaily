// Problem: https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/
// Results: https://leetcode.com/problems/minimum-number-of-moves-to-seat-everyone/submissions/1287386151
package main

import "slices"

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Time: O(n*log(n))
// Space: O(1)
func minMovesToSeat(seats []int, students []int) int {
	slices.Sort(seats)
	slices.Sort(students)
	moves := 0
	for i := range seats {
		moves += abs(students[i] - seats[i])
	}
	return moves
}
