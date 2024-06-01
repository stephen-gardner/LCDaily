package main

import "testing"

func test(t *testing.T, grid [][]int, expected int) {
	if res := uniquePathsIII(grid); res != expected {
		for _, row := range grid {
			t.Log(row)
		}
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestUniquePathsIII(t *testing.T) {
	test(t, [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, -1}}, 2)
	test(t, [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 2}}, 4)
	test(t, [][]int{
		{0, 1},
		{2, 0}}, 0)
}
