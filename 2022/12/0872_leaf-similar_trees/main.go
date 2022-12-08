// Problem: https://leetcode.com/problems/leaf-similar-trees/
// Results: https://leetcode.com/problems/leaf-similar-trees/submissions/856431361/
package main

import "github.com/stephen-gardner/LCDaily/test/tree"

type TreeNode = tree.TreeNode

func getSequence(res *[]byte, curr *TreeNode) {
	if curr == nil {
		return
	}
	if curr.Left == nil && curr.Right == nil {
		*res = append(*res, byte(curr.Val))
		return
	}
	getSequence(res, curr.Left)
	getSequence(res, curr.Right)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var s1, s2 []byte
	getSequence(&s1, root1)
	getSequence(&s2, root2)
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
