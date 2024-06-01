package main

import "testing"

func test(t *testing.T, s, expected string) {
	res := reverseWords(s)
	if res != expected {
		t.Log("   Input:", s)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestReverseWords(t *testing.T) {
	test(t, "Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc")
	test(t, "God Ding", "doG gniD")
}
