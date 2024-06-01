package main

import "testing"

func test(t *testing.T, trees, expected [][]int) {
	res := outerTrees(trees)
	fail := len(res) != len(expected)
	if !fail {
		present := map[[2]int]bool{}
		for _, t := range res {
			present[[2]int{t[0], t[1]}] = true
		}
		for _, t := range expected {
			if _, have := present[[2]int{t[0], t[1]}]; !have {
				fail = true
				break
			}
		}
	}
	if fail {
		t.Log("   Input:", trees)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestOuterTrees(t *testing.T) {
	test(
		t,
		[][]int{{1, 1}, {2, 2}, {2, 0}, {2, 4}, {3, 3}, {4, 2}},
		[][]int{{1, 1}, {2, 0}, {3, 3}, {2, 4}, {4, 2}},
	)
	test(
		t,
		[][]int{{1, 2}, {2, 2}, {4, 2}},
		[][]int{{4, 2}, {2, 2}, {1, 2}},
	)
	test(
		t,
		[][]int{{1, 2}, {2, 2}, {4, 2}, {5, 2}, {6, 2}, {7, 2}},
		[][]int{{4, 2}, {6, 2}, {2, 2}, {5, 2}, {1, 2}, {7, 2}},
	)
}
