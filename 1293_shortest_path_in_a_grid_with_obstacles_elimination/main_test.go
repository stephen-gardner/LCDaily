package main

import (
	"fmt"
	"strings"
	"testing"
)

func test(t *testing.T, grid [][]int, k, expected int) {
	if res := shortestPath(grid, k); res != expected {
		var sb strings.Builder
		for _, row := range grid {
			sb.WriteString(fmt.Sprintln(row))
		}
		t.Log("k =", k)
		t.Logf("Grid:\n%s", sb.String())
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestShortestPath(t *testing.T) {
	test(
		t,
		[][]int{{0, 0, 0}, {1, 1, 0}, {0, 0, 0}, {0, 1, 1}, {0, 0, 0}},
		1,
		6,
	)
	test(
		t,
		[][]int{{0, 1, 1}, {1, 1, 1}, {1, 0, 0}},
		1,
		-1,
	)
}
