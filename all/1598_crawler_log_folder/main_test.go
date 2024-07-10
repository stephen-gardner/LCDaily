package main

import "testing"

func TestMinOperations(t *testing.T) {
	test := func(logs []string, expected int) {
		if res := minOperations(logs); res != expected {
			t.Fatalf("logs: %v\n%d != %d (expected)", logs, res, expected)
		}
	}
	test([]string{"d1/", "d2/", "../", "d21/", "./"}, 2)
	test([]string{"d1/", "d2/", "./", "d3/", "../", "d31/"}, 3)
}
