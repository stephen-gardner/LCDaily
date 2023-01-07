package main

import "testing"

func test(t *testing.T, gas, cost []int, expected int) {
	if res := canCompleteCircuit(gas, cost); res != expected {
		t.Log(" Gas:", gas)
		t.Log("Cost:", cost)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestCanCompleteCircuit(t *testing.T) {
	test(t, []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3)
	test(t, []int{2, 3, 4}, []int{3, 4, 3}, -1)
	test(t, []int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}, 4)
}
