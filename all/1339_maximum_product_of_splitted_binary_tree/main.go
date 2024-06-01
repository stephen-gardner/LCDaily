// Problem: https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/
// Results: https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/submissions/857817337/
package main

import "github.com/stephen-gardner/LCDaily/test/tree"

type TreeNode = tree.TreeNode

const mod = 1e9 + 7

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sumTree(curr *TreeNode) int {
	if curr == nil {
		return 0
	}
	curr.Val += sumTree(curr.Left) + sumTree(curr.Right)
	return curr.Val
}

func findMaxProduct(maxProduct *int, curr *TreeNode, sum int) {
	if curr == nil {
		return
	}
	*maxProduct = max(*maxProduct, curr.Val*(sum-curr.Val))
	findMaxProduct(maxProduct, curr.Left, sum)
	findMaxProduct(maxProduct, curr.Right, sum)
}

func maxProduct(root *TreeNode) int {
	sum := sumTree(root)
	maxProduct := 0
	findMaxProduct(&maxProduct, root, sum)
	return maxProduct % mod
}
