package main

import "testing"

func test(t *testing.T, words []string, expected int) {
	if res := longestPalindrome(words); res != expected {
		t.Log("Words:", words)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestLongestPalindrome(t *testing.T) {
	test(t, []string{"lc", "cl", "gg"}, 6)
	test(t, []string{"ab", "ty", "yt", "lc", "cl", "ab"}, 8)
	test(t, []string{"cc", "ll", "xx"}, 2)
}
