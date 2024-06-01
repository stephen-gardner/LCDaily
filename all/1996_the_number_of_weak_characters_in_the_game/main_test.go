package main

import "testing"

func test(t *testing.T, properties [][]int, expected int) {
	res := numberOfWeakCharacters(properties)
	if res != expected {
		t.Log("Input:", properties)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestNumberOfWeakCharacters(t *testing.T) {
	test(t, [][]int{{5, 5}, {6, 3}, {3, 6}}, 0)
	test(t, [][]int{{2, 2}, {3, 3}}, 1)
	test(t, [][]int{{1, 5}, {10, 4}, {4, 3}}, 1)
	test(t, [][]int{{1, 1}, {2, 1}, {2, 2}, {1, 2}}, 1)
	test(t, [][]int{{7, 9}, {10, 7}, {6, 9}, {10, 4}, {7, 5}, {7, 10}}, 2)
}
