package main

import (
	"math"
	"testing"

	"github.com/stephen-gardner/LCDaily/test/tree"
)

const null = math.MaxInt

func test(t *testing.T, nums []int, expected int) {
	tree := tree.Create(nums, 0)
	if res := maxPathSum(tree); res != expected {
		t.Log("Input:", tree.String())
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestMaxPathSum(t *testing.T) {
	test(t, []int{1, 2, 3}, 6)
	test(t, []int{-10, 9, 20, null, null, 15, 7}, 42)
	test(t, []int{-3}, -3)
	test(t, []int{2, -1}, 2)
	test(t, []int{1, -2, 3}, 4)
	test(t, []int{
		9,
		6, -3,
		null, null, -6, 2,
		null, null, null, null, null, null, 2, null,
		null, null, null, null, null, null, null, null, null, null, null, null, -6, -6,
	}, 16)
}
