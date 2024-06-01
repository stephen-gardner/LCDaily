package main

import "testing"

func test(t *testing.T, n int, expected string) {
	if res := countAndSay(n); res != expected {
		t.Log("n =", n)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestCountAndSay(t *testing.T) {
	test(t, 1, "1")
	test(t, 4, "1211")
}
