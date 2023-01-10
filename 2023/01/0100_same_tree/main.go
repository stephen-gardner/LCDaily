// Problem: https://leetcode.com/problems/same-tree/
// Results: https://leetcode.com/problems/same-tree/submissions/875173418/
package main

import "github.com/stephen-gardner/LCDaily/test/tree"

type TreeNode = tree.TreeNode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
