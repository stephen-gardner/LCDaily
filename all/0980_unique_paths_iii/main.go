// Problem: https://leetcode.com/problems/unique-paths-iii/
// Results: https://leetcode.com/problems/unique-paths-iii/submissions/868395733/
package main

func findEnds(grid [][]int) (start, end [2]int, empty int) {
	height, width := len(grid), len(grid[0])
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch grid[y][x] {
			case 0:
				empty++
			case 1:
				start = [2]int{y, x}
			case 2:
				end = [2]int{y, x}
			}
		}
	}
	return
}

func explore(res *int, grid [][]int, empty, y, x, steps int) {
	switch grid[y][x] {
	case 2:
		if steps == empty {
			*res++
		}
		return
	case 0:
		steps++
		c := grid[y][x]
		grid[y][x] = -1
		if y-1 >= 0 {
			explore(res, grid, empty, y-1, x, steps)
		}
		if y+1 < len(grid) {
			explore(res, grid, empty, y+1, x, steps)
		}
		if x-1 >= 0 {
			explore(res, grid, empty, y, x-1, steps)
		}
		if x+1 < len(grid[0]) {
			explore(res, grid, empty, y, x+1, steps)
		}
		grid[y][x] = c
	}
}

func uniquePathsIII(grid [][]int) int {
	res := 0
	start, _, empty := findEnds(grid)
	grid[start[0]][start[1]] = 0
	explore(&res, grid, empty, start[0], start[1], -1)
	return res
}
