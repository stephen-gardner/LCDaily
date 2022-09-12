package main

import "testing"

func test(t *testing.T, speed []int, efficiency []int, n, k, expected int) {
	res := maxPerformance(n, speed, efficiency, k)
	if res != expected {
		t.Log("          k:", k)
		t.Log("      Speed:", speed)
		t.Log("Efficiency:", efficiency)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestMaxPerformance(t *testing.T) {
	test(t, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 6, 2, 60)
	test(t, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 6, 3, 68)
	test(t, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 6, 4, 72)
}
