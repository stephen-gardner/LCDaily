package main

import "testing"

func test(test *testing.T, s, t, expected string) {
	res := minWindow(s, t)
	if res != expected {
		test.Log("s:", s)
		test.Log("t:", t)
		test.Fatalf(`"%s" != "%s"`, res, expected)
	}
}

func TestMinWindow(t *testing.T) {
	test(t, "ADOBECODEBANC", "ABC", "BANC")
	test(t, "a", "a", "a")
	test(t, "a", "aa", "")
	test(t, "aa", "aa", "aa")
	test(t, "abc", "cba", "abc")
	test(t, "cabwefgewcwaefgcf", "cae", "cwae")
	test(t, "bba", "ab", "ba")
}
