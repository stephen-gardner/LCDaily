package main

import "testing"

func TestMinMovesToSeat(t *testing.T) {
	test := func(seats, students []int, expected int) {
		if res := minMovesToSeat(seats, students); res != expected {
			t.Log("seats:", seats)
			t.Log("students:", students)
			t.Fatalf("%d != %d (expected)", res, expected)
		}
	}
	test(
		[]int{3, 1, 5},
		[]int{2, 7, 4},
		4,
	)
	test(
		[]int{4, 1, 5, 9},
		[]int{1, 3, 2, 6},
		7,
	)
	test(
		[]int{2, 2, 6, 6},
		[]int{1, 3, 2, 6},
		4,
	)
	test(
		[]int{12, 14, 19, 19, 12},
		[]int{19, 2, 17, 20, 7},
		19,
	)
}
