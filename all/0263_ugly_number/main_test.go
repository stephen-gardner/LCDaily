package main

import "testing"

func test(t *testing.T, n int, expected bool) {
	if isUgly(n) != expected {
		t.Log("Input:", n)
		t.Fatalf("%v != %v (expected)", !expected, expected)
	}
}

func TestIsUgly(t *testing.T) {
	test(t, 6, true)
	test(t, 1, true)
	test(t, 14, false)
	test(t, 0, false)
	test(t, -2, false)
	test(t, 905391974, false)
}
