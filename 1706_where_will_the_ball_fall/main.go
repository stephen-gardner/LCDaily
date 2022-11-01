// Problem: https://leetcode.com/problems/where-will-the-ball-fall/
// Results: https://leetcode.com/submissions/detail/834386987/
package main

func findBall(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	res := make([]int, n)
	for x := 0; x < n; x++ {
		pos := x
		for y := 0; y < m; y++ {
			delta := grid[y][pos]
			pos += delta
			if pos == -1 || pos == n || grid[y][pos] != delta {
				res[x] = -1
				break
			}
		}
		if res[x] != -1 {
			res[x] = pos
		}
	}
	return res
}
