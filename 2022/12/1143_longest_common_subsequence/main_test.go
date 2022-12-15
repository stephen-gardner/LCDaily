package main

import "testing"

func test(t *testing.T, text1, text2 string, expected int) {
	if res := longestCommonSubsequence(text1, text2); res != expected {
		t.Log("Text1:", text1)
		t.Log("Text2:", text2)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestLongestCommonSubsequence(t *testing.T) {
	test(t, "abcde", "ace", 3)
	test(t, "abc", "abc", 3)
	test(t, "abc", "def", 0)
	test(t, "ezupkr", "ubmrapg", 2)
}
