package main

import "testing"

func TestFindComplement(t *testing.T) {
	test := func(num, expected int) {
		if res := findComplement(num); res != expected {
			t.Log("Num:", num)
			t.Fatalf("%d != %d (expected)", res, expected)
		}
	}
	test(5, 2)
	test(1, 0)
}
