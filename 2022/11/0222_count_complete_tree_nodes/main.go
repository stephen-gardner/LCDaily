// Problem: https://leetcode.com/problems/count-complete-tree-nodes/
// Results: https://leetcode.com/submissions/detail/843648554/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	total := 0
	root.Val = 0
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nLeft, nRight := uint16(curr.Val>>16), uint16(curr.Val&0xFFFF)
		if nLeft == 0 {
			for tmp := curr; tmp != nil; tmp = tmp.Left {
				nLeft++
			}
		}
		if nRight == 0 {
			for tmp := curr; tmp != nil; tmp = tmp.Right {
				nRight++
			}
		}
		if nLeft == nRight {
			total += (1 << nLeft) - 1
		} else {
			total++
			if nLeft > 1 {
				curr.Left.Val = (int(nLeft - 1)) << 16
				stack = append(stack, curr.Left)
			}
			if nRight > 1 {
				curr.Right.Val = int(nRight - 1)
				stack = append(stack, curr.Right)
			}
		}
	}
	return total
}
