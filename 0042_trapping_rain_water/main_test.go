package main

import "testing"

func test(t *testing.T, height []int, expected int) {
	res := trap(height)
	if res != expected {
		t.Log("Height:", height)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestTrap(t *testing.T) {
	test(t, []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6)
	test(t, []int{4, 2, 0, 3, 2, 5}, 9)
	test(t, []int{5, 4, 1, 2}, 1)
	test(t, []int{9, 6, 8, 8, 5, 6, 3}, 3)
	test(t, []int{8, 8, 1, 5, 6, 2, 5, 3, 3, 9}, 31)
	test(t, []int{9, 6, 7, 8, 4, 6, 8, 7, 2, 0, 1, 6, 2, 1, 9, 2, 8}, 65)
	test(t, []int{8, 9, 9, 5, 9, 1, 3, 5, 8, 0, 5, 2, 4, 3, 7, 1, 5, 5, 0, 0, 0, 5, 4, 2}, 59)
	test(t, []int{4, 0, 9, 6, 8, 4, 0, 0, 0, 7, 3, 8, 3, 2, 8, 9, 5, 6}, 64)
	test(t, []int{3, 1, 6, 5, 3, 9, 5, 9, 2, 2, 6, 8, 2, 6, 4, 8, 5, 0, 7, 3, 6, 1, 2, 6, 7, 4, 0, 7, 9}, 104)
}
