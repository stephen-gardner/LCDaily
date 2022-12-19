// Problem: https://leetcode.com/problems/find-if-path-exists-in-graph/
// Results: https://leetcode.com/problems/find-if-path-exists-in-graph/submissions/861877601/
package main

type Node struct {
	connected []*Node
}

func hasPath(nodes []*Node, src, dst *Node) bool {
	stack := []*Node{src}
	visited := map[*Node]bool{}
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[curr] {
			continue
		}
		visited[curr] = true
		if curr == dst {
			return true
		}
		for _, next := range curr.connected {
			stack = append(stack, next)
		}
	}
	return false
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	nodes := make([]*Node, n)
	for _, e := range edges {
		if nodes[e[0]] == nil {
			nodes[e[0]] = &Node{}
		}
		if nodes[e[1]] == nil {
			nodes[e[1]] = &Node{}
		}
		src, dst := nodes[e[0]], nodes[e[1]]
		src.connected = append(src.connected, dst)
		dst.connected = append(dst.connected, src)
	}
	return hasPath(nodes, nodes[source], nodes[destination])
}
