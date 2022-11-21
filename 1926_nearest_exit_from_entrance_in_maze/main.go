// Problem: https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/
// Results: https://leetcode.com/submissions/detail/847233055/
package main

type Location struct {
	y uint8
	x uint8
}

var dirs = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func nearestExit(maze [][]byte, entrance []int) int {
	endY, endX := uint8(len(maze))-1, uint8(len(maze[0]))-1
	queue := []Location{{uint8(entrance[0]), uint8(entrance[1])}}
	steps := 0
	for len(queue) > 0 {
		i, length := 0, len(queue)
		for i < length {
			pos := queue[i]
			if pos.y == 0 || pos.y == endY || pos.x == 0 || pos.x == endX {
				if pos.y != uint8(entrance[0]) || pos.x != uint8(entrance[1]) {
					return steps
				}
			}
			for _, dir := range dirs {
				y, x := pos.y+uint8(dir[0]), pos.x+uint8(dir[1])
				if y <= endY && x <= endX && maze[y][x] != '+' {
					maze[y][x] = '+'
					queue = append(queue, Location{y, x})
				}
			}
			i++
		}
		queue = queue[i:]
		steps++
	}
	return -1
}
