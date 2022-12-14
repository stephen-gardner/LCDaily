package main

import (
	"testing"
)

func test(t *testing.T, matrix [][]int, expected int) {
	if res := minFallingPathSum(matrix); res != expected {
		for _, row := range matrix {
			t.Log(row)
		}
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestMinFallingPathSum(t *testing.T) {
	test(t, [][]int{
		{2, 1, 3},
		{6, 5, 4},
		{7, 8, 9},
	}, 13)
	test(t, [][]int{
		{-19, 57},
		{-40, -5},
	}, -59)
	test(t, [][]int{
		{82, 81},
		{69, 33},
	}, 114)
}
