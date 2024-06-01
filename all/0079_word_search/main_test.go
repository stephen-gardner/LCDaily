package main

import "testing"

func test(t *testing.T, board [][]byte, word string, expected bool) {
	dupe := make([][]byte, len(board))
	for i := range dupe {
		dupe[i] = make([]byte, len(board[i]))
		for j := range dupe[i] {
			dupe[i][j] = board[i][j]
		}
	}
	if exist(dupe, word) != expected {
		for _, row := range board {
			t.Log("[", string(row), "]")
		}
		t.Log("Word:", word)
		t.Fatalf("%v != %v (expected)", !expected, expected)
	}
}

func TestExist(t *testing.T) {
	test(t, [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'}}, "ABCCED", true)
	test(t, [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'}}, "SEE", true)
	test(t, [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'}}, "ABCB", false)
}
