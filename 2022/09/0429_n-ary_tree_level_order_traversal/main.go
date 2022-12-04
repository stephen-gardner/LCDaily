// Problem: https://leetcode.com/problems/n-ary-tree-level-order-traversal/
// Results: https://leetcode.com/submissions/detail/791882661/
package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		level := make([]int, len(queue))
		nextQueue := []*Node{}
		for i := range queue {
			level[i] = queue[i].Val
			for _, child := range queue[i].Children {
				nextQueue = append(nextQueue, child)
			}
		}
		res = append(res, level)
		queue = nextQueue
	}
	return res
}
