package main

import "testing"

func test(t *testing.T, dominoes, expected string) {
	if res := pushDominoes(dominoes); res != expected {
		t.Log("   Input:", dominoes)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestPushDominoes(t *testing.T) {
	test(t, "RR.L", "RR.L")
	test(t, ".L.R...LR..L..", "LL.RR.LLRRLL..")
	test(t, "R.......L.R.........", "RRRR.LLLL.RRRRRRRRRR")
}
