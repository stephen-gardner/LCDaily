package main

import "testing"

func test(t *testing.T, s1, s2, baseStr, expected string) {
	if res := smallestEquivalentString(s1, s2, baseStr); res != expected {
		t.Log("     s1:", s1)
		t.Log("     s2:", s2)
		t.Log("baseStr:", baseStr)
		t.Fatalf(`"%s" != "%s" (expected)`, res, expected)
	}
}

func TestSmallestEquivalentString(t *testing.T) {
	test(t, "parker", "morris", "parser", "makkek")
	test(t, "hello", "world", "hold", "hdld")
	test(t, "leetcode", "programs", "sourcecode", "aauaaaaada")
}
