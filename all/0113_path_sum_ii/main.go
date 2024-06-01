// Problem: https://leetcode.com/problems/path-sum-ii/
// Results: https://leetcode.com/submissions/detail/807200510/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	var sumNodes func(*TreeNode, []int, int)
	sumNodes = func(curr *TreeNode, values []int, sum int) {
		if curr == nil {
			return
		}
		sum += curr.Val
		values = append(values, curr.Val)
		if curr.Left == curr.Right && sum == targetSum {
			path := make([]int, len(values))
			copy(path, values)
			res = append(res, path)
			return
		}
		sumNodes(curr.Left, values, sum)
		sumNodes(curr.Right, values, sum)
	}
	sumNodes(root, []int{}, 0)
	return res
}
