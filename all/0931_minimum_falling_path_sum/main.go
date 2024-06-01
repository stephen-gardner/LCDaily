// Problem: https://leetcode.com/problems/minimum-falling-path-sum/
// Results: https://leetcode.com/problems/minimum-falling-path-sum/submissions/859453685/
package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var ceil = 1 + int(1e4)
var dirs = [][]int{{-1, -1}, {-1, 0}, {-1, 1}}

func minFallingPathSum(matrix [][]int) int {
	height, width := len(matrix), len(matrix[0])
	for y := 1; y < height; y++ {
		for x := 0; x < width; x++ {
			curr, minSum := matrix[y][x], ceil
			for _, dir := range dirs {
				prevY, prevX := y+dir[0], x+dir[1]
				if prevY >= 0 && prevY < height && prevX >= 0 && prevX < width {
					minSum = min(minSum, curr+matrix[prevY][prevX])
				}
			}
			matrix[y][x] = minSum
		}
	}
	minSum := ceil
	for _, n := range matrix[len(matrix)-1] {
		minSum = min(minSum, n)
	}
	return minSum
}
