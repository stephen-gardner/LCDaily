// Problem: https://leetcode.com/problems/two-sum-iv-input-is-a-bst/
// Results: https://leetcode.com/submissions/detail/819002078/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	visited := map[int]bool{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curr == nil {
			continue
		}
		if _, seen := visited[k-curr.Val]; seen {
			return true
		}
		visited[curr.Val] = true
		stack = append(stack, curr.Left)
		stack = append(stack, curr.Right)
	}
	return false
}
