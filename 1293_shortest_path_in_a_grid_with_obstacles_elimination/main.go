// Problem: https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/
// Results: https://leetcode.com/submissions/detail/833251151/
package main

type Position struct {
	y, x, k, steps int16
}

var dirs = [4][2]int16{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func shortestPath(grid [][]int, k int) int {
	m, n := int16(len(grid)), int16(len(grid[0]))
	queue := []Position{{0, 0, int16(k), 0}}
	visited := map[int32]bool{}
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		if pos.y == m-1 && pos.x == n-1 {
			return int(pos.steps)
		}
		for _, d := range dirs {
			next := Position{pos.y + d[0], pos.x + d[1], pos.k, pos.steps + 1}
			if next.y < 0 || next.y >= m || next.x < 0 || next.x >= n {
				continue
			}
			key := (int32(next.k) << 16) | (int32(next.x) << 8) | int32(next.y)
			if visited[key] {
				continue
			}
			visited[key] = true
			if grid[next.y][next.x] == 1 {
				if next.k == 0 {
					continue
				}
				next.k--
			}
			queue = append(queue, next)
		}
	}
	return -1
}
