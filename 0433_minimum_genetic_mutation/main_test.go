package main

import "testing"

func test(t *testing.T, start, end string, bank []string, expected int) {
	if res := minMutation(start, end, bank); res != expected {
		t.Log("Start:", start)
		t.Log("  End:", end)
		t.Log(" Bank:", bank)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestMinMutation(t *testing.T) {
	test(
		t, "AACCGGTT", "AACCGGTA",
		[]string{"AACCGGTA"},
		1,
	)
	test(
		t, "AACCGGTT", "AAACGGTA",
		[]string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
		2,
	)
	test(
		t, "AAAAACCC", "AACCCCCC",
		[]string{"AAAACCCC", "AAACCCCC", "AACCCCCC"},
		3,
	)
}
