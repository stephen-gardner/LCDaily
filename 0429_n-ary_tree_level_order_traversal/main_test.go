package main

import (
	"math"
	"testing"
)

const null = math.MaxInt

func naryTreeCreate(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}
	children := make([][]int, 1)
	i := 0
	for _, n := range nums {
		if n == null {
			children = append(children, []int{})
			i++
		} else {
			children[i] = append(children[i], n)
		}
	}
	root := &Node{Val: children[0][0]}
	queue := []*Node{root}
	for i := 1; i < len(children); i++ {
		curr := queue[0]
		queue = queue[1:]
		childArr := children[i]
		for _, n := range childArr {
			node := &Node{Val: n}
			curr.Children = append(curr.Children, node)
			queue = append(queue, node)
		}
	}
	return root
}

func test(t *testing.T, nums []int, expected [][]int) {
	tree := naryTreeCreate(nums)
	res := levelOrder(tree)
	fail := len(res) != len(expected)
	if !fail {
		for i := 0; i < len(res); i++ {
			if fail = len(res[i]) != len(expected[i]); fail {
				break
			}
			for j := 0; j < len(res[i]); j++ {
				if fail = res[i][j] != expected[i][j]; fail {
					i = len(res)
					break
				}
			}
		}
	}
	if fail {
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestLevelOrder(t *testing.T) {
	test(
		t,
		[]int{1, null, 3, 2, 4, null, 5, 6},
		[][]int{{1}, {3, 2, 4}, {5, 6}},
	)
	test(
		t,
		[]int{0},
		[][]int{{0}},
	)
	test(
		t,
		[]int{1, null, 2, 3, 4, 5, null, null, 6, 7, null, 8, null, 9, 10, null, null, 11, null, 12, null, 13, null, null, 14},
		[][]int{{1}, {2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13}, {14}},
	)
	test(
		t,
		[]int{},
		[][]int{},
	)
}
