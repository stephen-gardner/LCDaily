package main

import "testing"

func test(t *testing.T, tasks [][]int, expected []int) {
	res := getOrder(tasks)
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
		t.Log("   Input:", tasks)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestGetOrder(t *testing.T) {
	test(t,
		[][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}},
		[]int{0, 2, 3, 1},
	)
	test(t,
		[][]int{{5, 2}, {7, 2}, {9, 4}, {6, 3}, {5, 10}, {1, 1}},
		[]int{5, 0, 1, 3, 2, 4})
	test(t,
		[][]int{{19, 13}, {16, 9}, {21, 10}, {32, 25}, {37, 4}, {49, 24}, {2, 15}, {38, 41}, {37, 34}, {33, 6}, {45, 4}, {18, 18}, {46, 39}, {12, 24}},
		[]int{6, 1, 2, 9, 4, 10, 0, 11, 5, 13, 3, 8, 12, 7})
}
