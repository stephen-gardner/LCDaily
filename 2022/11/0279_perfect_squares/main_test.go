package main

import "testing"

func test(t *testing.T, n, expected int) {
	if res := numSquares(n); res != expected {
		t.Fatalf("%d: %d != %d (expected)", n, res, expected)
	}
}

func TestNumSquares(t *testing.T) {
	test(t, 12, 3)
	test(t, 13, 2)
	test(t, 49, 1)
	test(t, 368, 4)
}
