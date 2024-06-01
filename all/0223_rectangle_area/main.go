// Problem: https://leetcode.com/problems/rectangle-area/
// Results: https://leetcode.com/submissions/detail/845054826/
package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	area := ((ax2 - ax1) * (ay2 - ay1)) + ((bx2 - bx1) * (by2 - by1))
	oLength := min(ax2, bx2) - max(ax1, bx1)
	oHeight := min(ay2, by2) - max(ay1, by1)
	if oLength > 0 && oHeight > 0 {
		area -= oLength * oHeight
	}
	return area
}
