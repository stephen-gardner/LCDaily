package main

import "testing"

func TestLemonadeChange(t *testing.T) {
	test := func(bills []int, expected bool) {
		if lemonadeChange(bills) != expected {
			t.Log("Bills:", bills)
			t.Fatalf("%v != %v (expected)", !expected, expected)
		}
	}
	test([]int{5, 5, 5, 10, 20}, true)
	test([]int{5, 5, 10, 10, 20}, false)
}
