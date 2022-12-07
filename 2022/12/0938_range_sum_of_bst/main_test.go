package main

import (
	"math"
	"testing"

	"github.com/stephen-gardner/LCDaily/test/tree"
)

const null = math.MaxInt

func test(t *testing.T, nums []int, low, high int) {
	expected := 0
	for _, n := range nums {
		if n >= low && n <= high {
			expected += n
		}
	}
	bst := tree.Create(nums, 0)
	if res := rangeSumBST(bst, low, high); res != expected {
		t.Log("Input:", bst.String())
		t.Logf("Range: [%d, %d]", low, high)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestRangeSumBST(t *testing.T) {
	test(t, []int{10, 5, 15, 3, 7, null, 18}, 7, 15)
	test(t, []int{10, 5, 15, 3, 7, 13, 18, 1, null, 6}, 6, 10)
	test(t, []int{0}, 1, 2)
	test(t, []int{1}, 1, 1)
}
