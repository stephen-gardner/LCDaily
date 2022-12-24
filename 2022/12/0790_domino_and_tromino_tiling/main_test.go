package main

import "testing"

func test(t *testing.T, n, expected int) {
	if res := numTilings(n); res != expected {
		t.Log("n =", n)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestNumTilings(t *testing.T) {
	test(t, 1, 1)
	test(t, 2, 2)
	test(t, 3, 5)
	test(t, 8, 258)
	test(t, 1000, 979232805)
}
