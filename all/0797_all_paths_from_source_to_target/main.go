// Problem: https://leetcode.com/problems/all-paths-from-source-to-target/
// Results: https://leetcode.com/problems/all-paths-from-source-to-target/submissions/867851305/
package main

type Position struct {
	index int
	path  []int
}

func allPathsSourceTarget(graph [][]int) [][]int {
	res := [][]int{}
	stack := []Position{{0, []int{0}}}
	dst := len(graph) - 1
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curr.index == dst {
			res = append(res, curr.path)
			continue
		}
		for i, next := range graph[curr.index] {
			path := curr.path
			if i != 0 {
				path = make([]int, len(curr.path))
				copy(path, curr.path)
			}
			path = append(path, next)
			stack = append(stack, Position{
				index: next,
				path:  path,
			})
		}
	}
	return res
}
