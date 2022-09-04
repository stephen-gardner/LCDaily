// Problem: https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/
// Results: https://leetcode.com/submissions/detail/791073216/
package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeWidth(root *TreeNode) (int, int) {
	left, right := 0, 0
	var measure func(*TreeNode, *int, *int, int)
	measure = func(curr *TreeNode, left, right *int, currWidth int) {
		if curr == nil {
			return
		}
		if currWidth > *right {
			*right = currWidth
		}
		if currWidth < *left {
			*left = currWidth
		}
		measure(curr.Left, left, right, currWidth-1)
		measure(curr.Right, left, right, currWidth+1)
	}
	measure(root, &left, &right, 0)
	return -left, right
}

func buildColumns(work [][]int, curr *TreeNode, level, pos int) {
	if curr == nil {
		return
	}
	// Shift level 10 bits to leave room for max value
	work[pos] = append(work[pos], level<<10|curr.Val)
	buildColumns(work, curr.Left, level+1, pos-1)
	buildColumns(work, curr.Right, level+1, pos+1)
}

func verticalTraversal(root *TreeNode) [][]int {
	left, right := treeWidth(root)
	res := make([][]int, 1+left+right)
	buildColumns(res, root, 1, left)
	for _, column := range res {
		sort.Ints(column)
		for i := 0; i < len(column); i++ {
			// Mask for value in last 10 bits
			column[i] &= 0x3FF
		}
	}
	return res
}
