package main

import "testing"

func test(t *testing.T, num int, expected string) {
	if res := intToRoman(num); res != expected {
		t.Log("     Num:", num)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestIntToRoman(t *testing.T) {
	test(t, 3, "III")
	test(t, 58, "LVIII")
	test(t, 1994, "MCMXCIV")
	test(t, 3999, "MMMCMXCIX")
}
