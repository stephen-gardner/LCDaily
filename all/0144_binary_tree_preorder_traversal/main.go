// Problem: https://leetcode.com/problems/binary-tree-preorder-traversal/
// Results: https://leetcode.com/problems/binary-tree-preorder-traversal/submissions/874420994/
package main

import "github.com/stephen-gardner/LCDaily/test/tree"

type TreeNode = tree.TreeNode

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	var buildResult func(*TreeNode)
	buildResult = func(curr *TreeNode) {
		res = append(res, curr.Val)
		if curr.Left != nil {
			buildResult(curr.Left)
		}
		if curr.Right != nil {
			buildResult(curr.Right)
		}
	}
	buildResult(root)
	return res
}
