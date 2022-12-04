// Problem: https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/
// Results: https://leetcode.com/submissions/detail/799436396/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {
	count := 0
	var countPalindromes func(*TreeNode, int)
	countPalindromes = func(curr *TreeNode, odd int) {
		if curr == nil {
			return
		}
		// 0 = even; 1 = odd
		odd ^= 1 << curr.Val
		// True only if both are nil
		if curr.Left == curr.Right {
			// Check if more than one bit is set
			if odd&(odd-1) == 0 {
				count++
			}
		}
		countPalindromes(curr.Left, odd)
		countPalindromes(curr.Right, odd)
	}
	countPalindromes(root, 0)
	return count
}
