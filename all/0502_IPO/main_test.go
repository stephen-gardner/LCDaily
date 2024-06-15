package main

import "testing"

func TestFindMaximizedCapital(t *testing.T) {
	test := func(k, w int, profits, capital []int, expected int) {
		if res := findMaximizedCapital(k, w, profits, capital); res != expected {
			t.Log("profits:", profits)
			t.Log("capital:", capital)
			t.Logf("k: %d\tw: %d", k, w)
			t.Fatalf("%d != %d (expected)", res, expected)
		}
	}
	test(2, 0, []int{1, 2, 3}, []int{0, 1, 1}, 4)
	test(3, 0, []int{1, 2, 3}, []int{0, 1, 2}, 6)
}
