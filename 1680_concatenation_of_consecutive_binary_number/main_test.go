package main

import "testing"

func test(t *testing.T, n, expected int) {
	if res := concatenatedBinary(n); res != expected {
		t.Log("   Input:", n)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestConcatenatedBinary(t *testing.T) {
	test(t, 1, 1)
	test(t, 3, 27)
	test(t, 12, 505379714)
	test(t, 42, 727837408)
	test(t, 100000, 757631812)
}
