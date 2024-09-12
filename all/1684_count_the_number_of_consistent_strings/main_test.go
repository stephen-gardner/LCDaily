package main

import "testing"

func TestCountConsistentStrings(t *testing.T) {
	test := func(allowed string, words []string, expected int) {
		if res := countConsistentStrings(allowed, words); res != expected {
			t.Log("Allowed:", allowed)
			t.Log("Words:", words)
			t.Fatalf("%d != %d (expected)", res, expected)
		}
	}
	test("ab", []string{"ad", "bd", "aaab", "baa", "badab"}, 2)
	test("abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}, 7)
	test("cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}, 4)
}
