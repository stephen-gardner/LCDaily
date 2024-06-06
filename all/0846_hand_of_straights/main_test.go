package main

import "testing"

func TestIsNStraightHand(t *testing.T) {
	test := func(hand []int, groupSize int, expected bool) {
		if isNStraightHand(hand, groupSize) != expected {
			t.Logf("Input: %v; groupSize = %d", hand, groupSize)
			t.Fatalf("%v != %v (expected)", !expected, expected)
		}
	}
	test([]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3, true)
	test([]int{1, 2, 3, 4, 5}, 4, false)
	test([]int{1, 1, 2, 2, 3, 3}, 3, true)
	test([]int{1, 5}, 1, true)
	test([]int{1, 5}, 2, false)
	test([]int{6, 5, 4, 3, 2, 1}, 3, true)
}
