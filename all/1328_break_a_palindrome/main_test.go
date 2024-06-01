package main

import "testing"

func test(t *testing.T, str, expected string) {
	if res := breakPalindrome(str); res != expected {
		t.Log("   Input:", str)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestBreakPalindrome(t *testing.T) {
	test(t, "abccba", "aaccba")
	test(t, "a", "")
	test(t, "zz", "az")
	test(t, "zzz", "azz")
	test(t, "zzzz", "azzz")
	test(t, "zzzzz", "azzzz")
	test(t, "zxz", "axz")
	test(t, "zxxz", "axxz")
	test(t, "zzxzz", "azxzz")
	test(t, "zxxxz", "axxxz")
	test(t, "zaaz", "aaaz")
	test(t, "aa", "ab")
	test(t, "aaa", "aab")
	test(t, "aaaa", "aaab")
	test(t, "aaaaa", "aaaab")
}
