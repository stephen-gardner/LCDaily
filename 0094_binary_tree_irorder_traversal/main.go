// Problem: https://leetcode.com/problems/binary-tree-inorder-traversal/
// Results: https://leetcode.com/submissions/detail/794309777/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var traverse func(*TreeNode)
	traverse = func(curr *TreeNode) {
		if curr == nil {
			return
		}
		traverse(curr.Left)
		res = append(res, curr.Val)
		traverse(curr.Right)
	}
	traverse(root)
	return res
}
