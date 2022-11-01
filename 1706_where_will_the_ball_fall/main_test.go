package main

import "testing"

func test(t *testing.T, grid [][]int, expected []int) {
	res := findBall(grid)
	fail := len(res) != len(expected)
	if !fail {
		for i := range expected {
			fail = res[i] != expected[i]
			if fail {
				break
			}
		}
	}
	if fail {
		for _, row := range grid {
			t.Log(row)
		}
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestFindBall(t *testing.T) {
	test(
		t,
		[][]int{
			{1, 1, 1, -1, -1},
			{1, 1, 1, -1, -1},
			{-1, -1, -1, 1, 1},
			{1, 1, 1, 1, -1},
			{-1, -1, -1, -1, -1},
		},
		[]int{1, -1, -1, -1, -1},
	)
	test(
		t,
		[][]int{{-1}},
		[]int{-1},
	)
	test(
		t,
		[][]int{
			{1, 1, 1, 1, 1, 1},
			{-1, -1, -1, -1, -1, -1},
			{1, 1, 1, 1, 1, 1},
			{-1, -1, -1, -1, -1, -1},
		},
		[]int{0, 1, 2, 3, 4, -1},
	)
}
