// Problem: https://leetcode.com/problems/range-sum-of-bst/
// Results: https://leetcode.com/problems/range-sum-of-bst/submissions/855852618/
package main

import "github.com/stephen-gardner/LCDaily/test/tree"

type TreeNode = tree.TreeNode

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	var sumNodes func(*TreeNode)
	sumNodes = func(curr *TreeNode) {
		if curr == nil {
			return
		}
		if curr.Val >= low && curr.Val <= high {
			sum += curr.Val
		}
		sumNodes(curr.Left)
		sumNodes(curr.Right)
	}
	sumNodes(root)
	return sum
}
