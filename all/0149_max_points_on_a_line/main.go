// Problem: https://leetcode.com/problems/max-points-on-a-line/
// Results: https://leetcode.com/problems/max-points-on-a-line/submissions/874561416/
package main

import (
	"math"
)

func getSlope(p1, p2 []int) float64 {
	m := float64(p2[1]-p1[1]) / float64(p2[0]-p1[0])
	if m == math.Inf(-1) {
		m = math.Inf(1)
	} else if m == 0 {
		m = 0
	}
	return m
}

func maxPoints(points [][]int) int {
	if len(points) == 1 {
		return 1
	}
	high := 0
	for i := 0; i < len(points); i++ {
		lines := map[float64]int{}
		for j := i + 1; j < len(points); j++ {
			eq := getSlope(points[i], points[j])
			if _, exists := lines[eq]; !exists {
				lines[eq] += 2
			} else {
				lines[eq]++
			}
		}
		for _, n := range lines {
			if n > high {
				high = n
			}
		}
	}
	return high
}
