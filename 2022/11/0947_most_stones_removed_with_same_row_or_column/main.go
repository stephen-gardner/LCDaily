// Problem: https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/
// Results: https://leetcode.com/submissions/detail/842990152/
package main

func removeStones(allStones [][]int) int {
	// Map coordinates to rows/columns in plane to index of stones in allStones
	rows := map[int][]int{}
	cols := map[int][]int{}
	for i, stone := range allStones {
		rows[stone[0]] = append(rows[stone[0]], i)
		cols[stone[1]] = append(cols[stone[1]], i)
	}
	removed := 0
	// Track stones already counted/removed
	visited := map[int]bool{}
	for _, row := range rows {
		if len(row) == 0 || visited[row[0]] {
			continue
		}
		// Stack up all stones that can be networked to current one and count
		//	the ones that have not yet been visited
		count := 0
		stack := []int{row[0]}
		for len(stack) > 0 {
			i := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if visited[i] {
				continue
			}
			visited[i] = true
			count++
			y, x := allStones[i][0], allStones[i][1]
			// All stones sharing row
			for _, n := range rows[y] {
				stack = append(stack, n)
			}
			rows[y] = nil
			// All stones sharing column
			for _, n := range cols[x] {
				stack = append(stack, n)
			}
			cols[x] = nil
		}
		if count > 0 {
			// Leave one stone
			removed += count - 1
		}
	}
	return removed
}
