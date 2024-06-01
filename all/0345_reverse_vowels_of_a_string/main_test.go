package main

import "testing"

func test(t *testing.T, s, expected string) {
	if res := reverseVowels(s); res != expected {
		t.Log("   Input:", s)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestReverseVowels(t *testing.T) {
	test(t, "hello", "holle")
	test(t, "leetcode", "leotcede")
	test(t, "aE", "Ea")
}
