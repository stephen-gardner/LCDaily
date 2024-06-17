package main

import "testing"

func TestJudgeSquareSum(t *testing.T) {
	test := func(c int, expected bool) {
		if judgeSquareSum(c) != expected {
			t.Fatalf("c = %d\n%v != %v (expected)", c, !expected, expected)
		}
	}
	test(0, true)
	test(1, true)
	test(2, true)
	test(3, false)
	test(5, true)
}
