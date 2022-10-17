package main

import "testing"

func test(t *testing.T, sentence string, expected bool) {
	if checkIfPangram(sentence) != expected {
		t.Log("Sentence:", sentence)
		t.Fatalf("%v != %v (expected)", !expected, expected)
	}
}

func TestCheckIfPangram(t *testing.T) {
	test(t, "thequickbrownfoxjumpsoverthelazydog", true)
	test(t, "leetcode", false)
}
