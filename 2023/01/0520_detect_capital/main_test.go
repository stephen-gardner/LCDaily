package main

import "testing"

func test(t *testing.T, word string, expected bool) {
	if detectCapitalUse(word) != expected {
		t.Log("Input:", word)
		t.Fatalf("%v != %v (expected)", !expected, expected)
	}
}

func TestDetectCapitalUse(t *testing.T) {
	test(t, "USA", true)
	test(t, "FlaG", false)
	test(t, "leetcode", true)
	test(t, "Google", true)
	test(t, "ffffffffffffffffffffF", false)
}
