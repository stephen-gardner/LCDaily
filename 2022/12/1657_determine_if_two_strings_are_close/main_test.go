package main

import "testing"

func test(t *testing.T, w1, w2 string, expected bool) {
	if closeStrings(w1, w2) != expected {
		t.Log("Word 1:", w1)
		t.Log("Word 2:", w2)
		t.Fatalf("%v != %v (expected)", !expected, expected)
	}
}

func TestCloseStrings(t *testing.T) {
	test(t, "abc", "bca", true)
	test(t, "a", "aa", false)
	test(t, "cabbba", "abbccc", true)
	test(t, "uau", "ssx", false)
}
