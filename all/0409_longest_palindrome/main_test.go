package main

import "testing"

func TestLongestPalindrome(t *testing.T) {
	test := func(input string, expected int) {
		if res := longestPalindrome(input); res != expected {
			t.Fatalf("\nInput: %s\n%d != %d (expected)", input, res, expected)
		}
	}
	test("abccccdd", 7)
	test("a", 1)
}
