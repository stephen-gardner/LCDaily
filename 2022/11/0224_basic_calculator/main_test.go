package main

import "testing"

func test(t *testing.T, s string, expected int) {
	if res := calculate(s); res != expected {
		t.Log("Input:", s)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestCalculate(t *testing.T) {
	test(t, "1 + 1", 2)
	test(t, " 2-1 + 2 ", 3)
	test(t, "(1+(4+5+2)-3)+(6+8)", 23)
	test(t, "4  +   ( (3 - 2  )+      1)", 6)
	test(t, "-(4-2)", -2)
	test(t, "-(2+3)", -5)
	test(t, "-(3+2-(5 +4 + 3))", 7)
	test(t, "-(((4 + 2) - 3) + 1)", -4)
	test(t, "-(-(4))", 4)
	test(t, "1-(     -2)", 3)
}
