package main

import "testing"

func test(t *testing.T, ax1, ay1, ax2, ay2, bx1, by1, bx2, by2, expected int) {
	if res := computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2); res != expected {
		t.Logf("\tax1: %3d, ay1: %3d", ax1, ay1)
		t.Logf("\tax2: %3d, ay2: %3d", ax2, ay2)
		t.Logf("\tbx1: %3d, by1: %3d", bx1, by1)
		t.Logf("\tbx2: %3d, by2: %3d", bx2, by2)
		t.Fatalf("\t%d != %d (expected)", res, expected)
	}
}

func TestComputeArea(t *testing.T) {
	test(t, -3, 0, 3, 4, 0, -1, 9, 2, 45)
	test(t, -2, -2, 2, 2, -2, -2, 2, 2, 16)
}
