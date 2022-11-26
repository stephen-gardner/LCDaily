package main

import "testing"

func test(t *testing.T, startTime, endTime, profit []int, expected int) {
	if res := jobScheduling(startTime, endTime, profit); res != expected {
		t.Log(" Start:", startTime)
		t.Log("   End:", endTime)
		t.Log("Profit:", profit)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestJobScheduling(t *testing.T) {
	test(t,
		[]int{1, 2, 3, 3},
		[]int{3, 4, 5, 6},
		[]int{50, 10, 40, 70}, 120)
	test(t,
		[]int{1, 2, 3, 4, 6},
		[]int{3, 5, 10, 6, 9},
		[]int{20, 20, 100, 70, 60}, 150)
	test(t,
		[]int{1, 1, 1},
		[]int{2, 3, 4},
		[]int{5, 6, 4}, 6)
	test(t,
		[]int{4, 2, 4, 8, 2},
		[]int{5, 5, 5, 10, 8},
		[]int{1, 2, 8, 10, 4}, 18)
}
