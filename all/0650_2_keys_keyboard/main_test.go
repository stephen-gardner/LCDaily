package main

import "testing"

func TestMinSteps(t *testing.T) {
	test := func(n, expected int) {
		if res := minSteps(n); res != expected {
			t.Log("n =", n)
			t.Fatalf("%d != %d (expected)", res, expected)
		}
	}
	test(1, 0)
	test(2, 2)
	test(3, 3)
	test(4, 4)
	test(5, 5)
	test(6, 5)
	test(12, 7)
	test(1000, 21)
}
