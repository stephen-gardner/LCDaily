package main

import "testing"

func test(t *testing.T, s string, expected int) {
	if res := numDecodings(s); res != expected {
		t.Log("Input:", s)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestNumDecodings(t *testing.T) {
	test(t, "11106", 2)
	test(t, "12", 2)
	test(t, "226", 3)
}

/*
1 11 06
1 1 10 6
11 10 6
*/
