// Problem: https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/
// Results: https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/submissions/871638330/
package main

import "sort"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMinArrowShots(points [][]int) int {
	arrows := 0
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	for i := 0; i < len(points); arrows++ {
		for endX := points[i][1]; i < len(points) && points[i][0] <= endX; i++ {
			endX = min(endX, points[i][1])
		}
	}
	return arrows
}
