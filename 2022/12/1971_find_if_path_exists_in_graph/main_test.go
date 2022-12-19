package main

import "testing"

func test(t *testing.T, edges [][]int, n, src, dst int, expected bool) {
	if validPath(n, edges, src, dst) != expected {
		t.Logf("      Edges: %v; n = %d", edges, n)
		t.Log("     Source:", src)
		t.Log("Destination:", dst)
		t.Fatalf("%v != %v (expected)", !expected, expected)
	}
}

func TestValidPath(t *testing.T) {
	test(t, [][]int{
		{0, 1}, {1, 2}, {2, 0}},
		3, 0, 2, true)
	test(t, [][]int{
		{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}},
		6, 0, 5, false)
}
