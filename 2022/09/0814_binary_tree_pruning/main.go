// Problem: https://leetcode.com/problems/binary-tree-pruning/
// Results: https://leetcode.com/submissions/detail/792791696/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	var prune func(*TreeNode) bool
	prune = func(curr *TreeNode) bool {
		if curr == nil {
			return false
		}
		if !prune(curr.Left) {
			curr.Left = nil
		}
		if !prune(curr.Right) {
			curr.Right = nil
		}
		return !(curr.Left == nil && curr.Right == nil && curr.Val == 0)
	}
	if !prune(root) {
		return nil
	}
	return root
}
