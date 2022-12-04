package main

import "testing"

func test(t *testing.T, s, expected string) {
	if res := makeGood(s); res != expected {
		t.Log("   Input:", s)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestMakeGood(t *testing.T) {
	test(t, "leEeetcode", "leetcode")
	test(t, "abBAcC", "")
	test(t, "s", "s")
}
