// Problem: https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/
// Results: https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/submissions/856900063/
package main

import "github.com/stephen-gardner/LCDaily/test/tree"

type TreeNode = tree.TreeNode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// O(n^2) solution would be to try to compute diff for the children of each
//	node, but since we want the max, we can make do with the extreme ends
//	(min/max) to calculate it in O(n) time.
func findMaxDiff(res *int, curr *TreeNode, minAncestor, maxAncestor int) {
	if curr == nil {
		return
	}
	minAncestor = min(minAncestor, curr.Val)
	maxAncestor = max(maxAncestor, curr.Val)
	*res = max(*res, maxAncestor-minAncestor)
	findMaxDiff(res, curr.Left, minAncestor, maxAncestor)
	findMaxDiff(res, curr.Right, minAncestor, maxAncestor)
}

func maxAncestorDiff(root *TreeNode) int {
	res := 0
	findMaxDiff(&res, root, root.Val, root.Val)
	return res
}
