package main

import "testing"

func test(t *testing.T, graph, expected [][]int) {
	res := allPathsSourceTarget(graph)
	fail := len(res) != len(expected)
	if !fail {
		for i := range expected {
			fail = len(res[i]) != len(expected[i])
			if fail {
				break
			}
			for j := range expected[i] {
				if res[i][j] != expected[i][j] {
					fail = true
					i = len(expected)
					break
				}
			}
		}
	}
	if fail {
		t.Log("   Input:", graph)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestAllPathsSourceTarget(t *testing.T) {
	test(t,
		[][]int{{1, 2}, {3}, {3}, {}},
		[][]int{{0, 2, 3}, {0, 1, 3}})
	test(t,
		[][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}},
		[][]int{{0, 1, 4}, {0, 1, 2, 3, 4}, {0, 1, 3, 4}, {0, 3, 4}, {0, 4}})
}
