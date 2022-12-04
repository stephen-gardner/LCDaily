// Problem: https://leetcode.com/problems/construct-string-from-binary-tree/
// Results: https://leetcode.com/submissions/detail/793561820/
package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	var sb strings.Builder
	var traverse func(*TreeNode)
	traverse = func(curr *TreeNode) {
		if curr == nil {
			return
		}
		sb.WriteString(fmt.Sprintf("%d", curr.Val))
		if curr.Left != nil || curr.Right != nil {
			sb.WriteByte('(')
			traverse(curr.Left)
			sb.WriteByte(')')
			if curr.Right != nil {
				sb.WriteByte('(')
				traverse(curr.Right)
				sb.WriteByte(')')
			}
		}
	}
	traverse(root)
	return sb.String()
}
