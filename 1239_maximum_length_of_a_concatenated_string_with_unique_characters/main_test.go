package main

import "testing"

func test(t *testing.T, arr []string, expected int) {
	dupe := make([]string, len(arr))
	for i := range arr {
		dupe[i] = arr[i]
	}
	if res := maxLength(dupe); res != expected {
		t.Log("Input:", arr)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestMaxLength(t *testing.T) {
	test(t, []string{"un", "iq", "ue"}, 4)
	test(t, []string{"cha", "r", "act", "ers"}, 6)
	test(t, []string{"abcdefghijklmnopqrstuvwxyz"}, 26)
	test(t, []string{"aa", "bb"}, 0)
}
