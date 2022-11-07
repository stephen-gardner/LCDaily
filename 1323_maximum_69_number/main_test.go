package main

import "testing"

func test(t *testing.T, num, expected int) {
	if res := maximum69Number(num); res != expected {
		t.Log("     Num:", num)
		t.Log("  Result:", res)
		t.Log("Expecetd:", expected)
		t.FailNow()
	}
}

func TestMaximum69Number(t *testing.T) {
	test(t, 9669, 9969)
	test(t, 9996, 9999)
	test(t, 9999, 9999)
}
