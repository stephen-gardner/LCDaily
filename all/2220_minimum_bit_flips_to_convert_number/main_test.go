package main

import "testing"

func TestMinBitFlips(t *testing.T) {
	test := func(start, goal, expected int) {
		if res := minBitFlips(start, goal); res != expected {
			t.Fatalf("Start: %d\nGoal: %d\nResult: %d != %d (expected)", start, goal, res, expected)
		}
	}
	test(10, 7, 3)
	test(3, 4, 3)
}
