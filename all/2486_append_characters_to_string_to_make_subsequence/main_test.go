package main

import "testing"

func TestAppendCharacters(t *testing.T) {
	test := func(ts *testing.T, s, t string, expected int) {
		if res := appendCharacters(s, t); res != expected {
			ts.Fatalf("\ns: %s\nt: %s\n%d != %d (expected)", s, t, res, expected)
		}
	}
	test(t, "coaching", "coding", 4)
	test(t, "abcde", "a", 0)
	test(t, "z", "abcde", 5)
}
