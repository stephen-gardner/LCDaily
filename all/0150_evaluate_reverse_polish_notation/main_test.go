package main

import "testing"

func test(t *testing.T, tokens []string, expected int) {
	if res := evalRPN(tokens); res != expected {
		t.Log("Input:", tokens)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestEvalRPN(t *testing.T) {
	test(t, []string{"2", "1", "+", "3", "*"}, 9)
	test(t, []string{"4", "13", "5", "/", "+"}, 6)
	test(t, []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22)
}
