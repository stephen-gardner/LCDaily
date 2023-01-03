package main

import "testing"

func test(t *testing.T, strs []string, expected int) {
	if res := minDeletionSize(strs); res != expected {
		t.Log("Input:", strs)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestMinDeletionSize(t *testing.T) {
	test(t, []string{"cba", "daf", "ghi"}, 1)
	test(t, []string{"a", "b"}, 0)
	test(t, []string{"zyx", "wvu", "tsr"}, 3)
	test(t, []string{"aaa", "aba", "aaa"}, 1)
}
