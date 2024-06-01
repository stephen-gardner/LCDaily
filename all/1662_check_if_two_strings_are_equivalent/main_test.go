package main

import "testing"

func test(t *testing.T, word1, word2 []string, expected bool) {
	if arrayStringsAreEqual(word1, word2) != expected {
		t.Log("Word1:", word1)
		t.Log("Word2:", word2)
		t.Fatalf("%v != %v (expected)", !expected, expected)
	}
}

func TestArrayStringsAreEqual(t *testing.T) {
	test(
		t,
		[]string{"ab", "c"},
		[]string{"a", "bc"},
		true,
	)
	test(
		t,
		[]string{"a", "cb"},
		[]string{"ab", "c"},
		false,
	)
	test(
		t,
		[]string{"abc", "d", "defg"},
		[]string{"abcddefg"},
		true,
	)
}
