package main

import "testing"

func test(t *testing.T, input string, expected int) {
	if res := scoreOfString(input); res != expected {
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestScoreOfString(t *testing.T) {
	test(t, "hello", 13)
	test(t, "zaz", 50)
}
