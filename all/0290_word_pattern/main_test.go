package main

import "testing"

func test(t *testing.T, pattern, s string, expected bool) {
	if wordPattern(pattern, s) != expected {
		t.Log(" String:", s)
		t.Log("Pattern:", pattern)
		t.Fatalf("%v != %v (expected)", !expected, expected)
	}
}

func TestWordPattern(t *testing.T) {
	test(t, "abba", "dog cat cat dog", true)
	test(t, "abba", "dog cat cat fish", false)
	test(t, "aaaa", "dog cat cat dog", false)
	test(t, "abba", "dog dog dog dog", false)
}
