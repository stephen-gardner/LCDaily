package main

import "testing"

func TestReverseParentheses(t *testing.T) {
	test := func(input, expected string) {
		if res := reverseParentheses(input); res != expected {
			t.Fatalf("\nInput: %s\nResult: %s\nExpected: %s", input, res, expected)
		}
	}
	test("(abcd)", "dcba")
	test("(u(love)i)", "iloveu")
	test("(ed(et(oc))el)", "leetcode")
	test("a(bcdefghijkl(mno)p)q", "apmnolkjihgfedcbq")
	test("sxmdll(q)eki(x)", "sxmdllqekix")
	test("vdgzyj()", "vdgzyj")
	test("((eqk((h))))", "eqkh")
}
