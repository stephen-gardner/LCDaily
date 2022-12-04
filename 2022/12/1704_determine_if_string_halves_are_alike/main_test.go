package main

import "testing"

func test(t *testing.T, s string, expected bool) {
	if halvesAreAlike(s) != expected {
		t.Log("Input:", s)
		t.Fatalf("%v != %v (expected)", !expected, expected)
	}
}

func TestHalvesAreAlike(t *testing.T) {
	test(t, "book", true)
	test(t, "textbook", false)
}
