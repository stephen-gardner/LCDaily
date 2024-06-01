package main

import "testing"

func test(t *testing.T, n, expected int) {
	if res := climbStairs(n); res != expected {
		t.Log("n =", n)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestClimbStairs(t *testing.T) {
	test(t, 2, 2)
	test(t, 3, 3)
	test(t, 50, 20365011074)
}
