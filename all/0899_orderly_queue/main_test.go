package main

import "testing"

func test(t *testing.T, k int, s, expected string) {
	if res := orderlyQueue(s, k); res != expected {
		t.Log("k =", k)
		t.Log("    Input:", s)
		t.Log("   Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestOrderlyQueue(t *testing.T) {
	test(t, 1, "cba", "acb")
	test(t, 3, "baaca", "aaabc")
}
