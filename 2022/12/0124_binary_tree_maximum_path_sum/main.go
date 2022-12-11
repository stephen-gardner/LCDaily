// Problem: https://leetcode.com/problems/binary-tree-maximum-path-sum/
// Results: https://leetcode.com/problems/binary-tree-maximum-path-sum/submissions/857864534/
package main

import "github.com/stephen-gardner/LCDaily/test/tree"

type TreeNode = tree.TreeNode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxSum(sum *int, curr *TreeNode) int {
	if curr == nil {
		return 0
	}
	left := findMaxSum(sum, curr.Left)
	right := findMaxSum(sum, curr.Right)
	// Options:
	//	- Current
	//	- Current + left path
	//	- Current + right path
	//	- Current + left path + right path
	bestPath := max(curr.Val, curr.Val+max(left, right))
	*sum = max(*sum, max(bestPath, curr.Val+left+right))
	return bestPath
}

func maxPathSum(root *TreeNode) int {
	sum := -1000
	findMaxSum(&sum, root)
	return sum
}
