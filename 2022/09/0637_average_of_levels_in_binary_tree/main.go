// Problem: https://leetcode.com/problems/average-of-levels-in-binary-tree/
// Results: https://leetcode.com/submissions/detail/789468648/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumLevels(totals map[int][2]int, curr *TreeNode, level int) {
	if curr == nil {
		return
	}
	totals[level] = [2]int{totals[level][0] + curr.Val, totals[level][1] + 1}
	sumLevels(totals, curr.Left, level+1)
	sumLevels(totals, curr.Right, level+1)
}

func averageOfLevels(root *TreeNode) []float64 {
	totals := map[int][2]int{}
	sumLevels(totals, root, 0)
	res := make([]float64, len(totals))
	for level, sum := range totals {
		res[level] = float64(sum[0]) / float64(sum[1])
	}
	return res
}
