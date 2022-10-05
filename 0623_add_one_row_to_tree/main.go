// Problem: https://leetcode.com/problems/add-one-row-to-tree/
// Results: https://leetcode.com/submissions/detail/815557615/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addRow(curr *TreeNode, val int, depth int) *TreeNode {
	if curr == nil {
		return nil
	}
	if depth == 2 {
		curr.Left = &TreeNode{
			Val:  val,
			Left: curr.Left,
		}
		curr.Right = &TreeNode{
			Val:   val,
			Right: curr.Right,
		}
		return curr
	}
	depth--
	addRow(curr.Left, val, depth)
	addRow(curr.Right, val, depth)
	return curr
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{
			Val:  val,
			Left: root,
		}
	}
	return addRow(root, val, depth)
}
