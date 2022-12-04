package main

import "testing"

func test(t *testing.T, tokens []int, power, expected int) {
	res := bagOfTokensScore(tokens, power)
	if res != expected {
		t.Log("Tokens:", tokens)
		t.Log(" Power:", power)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestBagOfTokensScore(t *testing.T) {
	test(t, []int{100}, 50, 0)
	test(t, []int{100, 200}, 150, 1)
	test(t, []int{100, 200, 300, 400}, 200, 2)
}
