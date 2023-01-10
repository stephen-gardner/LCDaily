package main

import (
	"math"
	"testing"

	"github.com/stephen-gardner/LCDaily/test/tree"
)

const null = math.MaxInt

func test(t *testing.T, nums1, nums2 []int, expected bool) {
	t1, t2 := tree.Create(nums1, 0), tree.Create(nums2, 0)
	if isSameTree(t1, t2) != expected {
		t.Log("p:", t1)
		t.Log("q:", t2)
		t.Fatalf("%v != %v (expected)", !expected, expected)
	}
}

func TestIsSameTree(t *testing.T) {
	test(t, []int{1, 2, 3}, []int{1, 2, 3}, true)
	test(t, []int{1, 2}, []int{1, null, 2}, false)
	test(t, []int{1, 2, 1}, []int{1, 1, 2}, false)
}
