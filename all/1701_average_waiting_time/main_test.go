package main

import "testing"

func TestAverageWaitingTimet(t *testing.T) {
	test := func(customers [][]int, expected float64) {
		if res := averageWaitingTime(customers); res != expected {
			t.Fatalf("customers: %v\n%f != %f (expected)", customers, res, expected)
		}
	}
	test([][]int{{1, 2}, {2, 5}, {4, 3}}, 5.0)
	test([][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}}, 3.25)
}
