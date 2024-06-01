package main

import "testing"

func test(t *testing.T, edges [][]int, labels string, expected []int) {
	res := countSubTrees(len(expected), edges, labels)
	fail := len(res) != len(expected)
	if !fail {
		for i := range expected {
			if res[i] != expected[i] {
				fail = true
				break
			}
		}
	}
	if fail {
		t.Log("n =", len(expected))
		t.Log("  Labels:", labels)
		t.Log("   Edges:", edges)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestCountSubTrees(t *testing.T) {
	test(t, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, "abaedcd", []int{2, 1, 1, 1, 1, 1, 1})
	test(t, [][]int{{0, 1}, {1, 2}, {0, 3}}, "bbbb", []int{4, 2, 1, 1})
	test(t, [][]int{{0, 1}, {0, 2}, {1, 3}, {0, 4}}, "aabab", []int{3, 2, 1, 1, 1})
}
