package main

import "testing"

func TestPassThePillow(t *testing.T) {
	test := func(n, time, expected int) {
		if res := passThePillow(n, time); res != expected {
			t.Fatalf("n = %d; time = %d\n%d != %d (expected)", n, time, res, expected)
		}
	}
	test(4, 5, 2)
	test(3, 2, 3)
	test(5, 0, 1)
	test(5, 1, 2)
	test(5, 4, 5)
	test(18, 38, 5)
}
