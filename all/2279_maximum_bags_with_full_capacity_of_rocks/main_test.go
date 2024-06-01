package main

import "testing"

func test(t *testing.T, capacity, rocks []int, additionalRocks, expected int) {
	if res := maximumBags(capacity, rocks, additionalRocks); res != expected {
		t.Log("  Capacity:", capacity)
		t.Log("     Rocks:", rocks)
		t.Log("Additional:", additionalRocks)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestMaximumBags(t *testing.T) {
	test(t,
		[]int{2, 3, 4, 5},
		[]int{1, 2, 4, 4},
		2, 3)
	test(t,
		[]int{10, 2, 2},
		[]int{2, 2, 0},
		100, 3)
	test(t,
		[]int{54, 18, 91, 49, 51, 45, 58, 54, 47, 91, 90, 20, 85, 20, 90, 49, 10, 84, 59, 29, 40, 9, 100, 1, 64, 71, 30, 46, 91},
		[]int{14, 13, 16, 44, 8, 20, 51, 15, 46, 76, 51, 20, 77, 13, 14, 35, 6, 34, 34, 13, 3, 8, 1, 1, 61, 5, 2, 15, 18},
		77, 13)
}
