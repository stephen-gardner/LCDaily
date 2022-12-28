package main

import "testing"

func test(t *testing.T, piles []int, k, expected int) {
	if res := minStoneSum(piles, k); res != expected {
		t.Log("Input:", piles)
		t.Log("k =", k)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestMinStoneSum(t *testing.T) {
	test(t, []int{5, 4, 9}, 2, 12)
	test(t, []int{4, 3, 6, 7}, 3, 12)
	test(t, []int{2695, 9184, 2908, 3869, 3779, 391, 2896, 5328}, 10, 10946)
}
