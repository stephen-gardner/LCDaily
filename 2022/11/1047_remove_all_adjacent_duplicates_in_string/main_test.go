package main

import "testing"

func test(t *testing.T, s, expected string) {
	if res := removeDuplicates(s); res != expected {
		t.Log("   Input:", s)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestRemoveDuplicates(t *testing.T) {
	test(t, "abbaca", "ca")
	test(t, "azxxzy", "ay")
}
