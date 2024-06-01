// Problem: https://leetcode.com/problems/longest-path-with-different-adjacent-characters/
// Results: https://leetcode.com/problems/longest-path-with-different-adjacent-characters/submissions/878296173/
package main

type Node struct {
	connected []*Node
	d1        int
	d2        int
}

// Fan out in all directions and record two longest branches
func measurePaths(prev, curr *Node) {
	curr.d1 = 1
	for _, next := range curr.connected {
		if next == prev {
			continue
		}
		if next.d1 == 0 {
			measurePaths(curr, next)
		}
		if length := 1 + next.d1; length > curr.d2 {
			curr.d2 = length
		}
		if curr.d2 > curr.d1 {
			curr.d1, curr.d2 = curr.d2, curr.d1
		}
	}
}

func longestPath(parent []int, s string) int {
	nodes := make([]*Node, len(parent))
	for i := range nodes {
		nodes[i] = &Node{d2: 1}
	}
	for i := 1; i < len(parent); i++ {
		// Don't link nodes that have the same character, as we can't travel
		//	down them
		p := parent[i]
		if s[i] != s[p] {
			n1, n2 := nodes[i], nodes[p]
			n1.connected = append(n1.connected, n2)
			n2.connected = append(n2.connected, n1)
		}
	}
	longest := 0
	for _, curr := range nodes {
		if curr.d1 == 0 {
			measurePaths(nil, curr)
		}
		// -1 because each branch length includes the current node
		if length := curr.d1 + curr.d2 - 1; length > longest {
			longest = length
		}
	}
	return longest
}
