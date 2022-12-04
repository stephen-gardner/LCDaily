package main

import "testing"

func test(t *testing.T, maze [][]byte, entrance []int, expected int) {
	dupe := make([][]byte, len(maze))
	for i := range dupe {
		dupe[i] = make([]byte, len(maze[i]))
		for j := range dupe[i] {
			dupe[i][j] = maze[i][j]
		}
	}
	if res := nearestExit(dupe, entrance); res != expected {
		for _, row := range maze {
			t.Log(string(row))
		}
		t.Log("Entrance:", entrance)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestNearestExit(t *testing.T) {
	test(
		t,
		[][]byte{
			{'+', '+', '.', '+'},
			{'.', '.', '.', '+'},
			{'+', '+', '+', '.'},
		},
		[]int{1, 2},
		1,
	)
	test(
		t,
		[][]byte{
			{'+', '+', '+'},
			{'.', '.', '.'},
			{'+', '+', '+'},
		},
		[]int{1, 0},
		2,
	)
	test(
		t,
		[][]byte{
			{'.', '+'},
		},
		[]int{0, 0},
		-1,
	)
}
