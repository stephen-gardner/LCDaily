package main

import "testing"

func test(t *testing.T, s, expected string) {
	if res := reverseWords(s); res != expected {
		t.Logf("   Input: \"%s\"", s)
		t.Logf("  Result: \"%s\"", res)
		t.Logf("Expected: \"%s\"", expected)
		t.FailNow()
	}
}

func TestReverseWords(t *testing.T) {
	test(t, "the sky is blue", "blue is sky the")
	test(t, "  hello world  ", "world hello")
	test(t, "a good   example", "example good a")
}
