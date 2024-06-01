// Problem: https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/
// Results: https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/submissions/876632211/
package main

type Node struct {
	id        int
	connected []*Node
	label     byte
}

func count(res []int, curr, prev *Node) [26]int {
	var counts [26]int
	if curr == nil {
		return counts
	}
	counts[curr.label]++
	for _, next := range curr.connected {
		if next != prev {
			for i, n := range count(res, next, curr) {
				counts[i] += n
			}
		}
	}
	res[curr.id] = counts[curr.label]
	return counts
}

func countSubTrees(n int, edges [][]int, labels string) []int {
	nodes := map[int]*Node{}
	for i := 0; i < n; i++ {
		nodes[i] = &Node{
			id:    i,
			label: labels[i] - 'a',
		}
	}
	for _, e := range edges {
		n1, n2 := nodes[e[0]], nodes[e[1]]
		n1.connected = append(n1.connected, n2)
		n2.connected = append(n2.connected, n1)
	}
	res := make([]int, n)
	count(res, nodes[0], nil)
	return res
}
