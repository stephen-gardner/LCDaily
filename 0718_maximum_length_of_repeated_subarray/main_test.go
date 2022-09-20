package main

import "testing"

func test(t *testing.T, nums1, nums2 []int, expected int) {
	res := findLength(nums1, nums2)
	if res != expected {
		t.Log("Nums1:", nums1)
		t.Log("Nums2:", nums2)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestFindLength(t *testing.T) {
	test(t, []int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}, 3)
	test(t, []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, 5)
	test(t, []int{1, 0, 1, 0, 0, 0, 0, 0, 1, 1}, []int{1, 1, 0, 1, 1, 0, 0, 0, 0, 0}, 6)
}
