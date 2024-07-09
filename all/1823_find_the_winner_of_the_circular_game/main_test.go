package main

import "testing"

func TestFindTheWinner(t *testing.T) {
	test := func(n, k, expected int) {
		if res := findTheWinner(n, k); res != expected {
			t.Fatalf("n = %d, k = %d\n%d != %d (expecetd)", n, k, res, expected)
		}
	}
	test(5, 2, 3)
	test(6, 5, 1)
}
