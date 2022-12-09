package main

import (
	"math"
	"testing"

	"github.com/stephen-gardner/LCDaily/test/tree"
)

const null = math.MaxInt

func test(t *testing.T, nums []int, expected int) {
	tree := tree.Create(nums, 0)
	if res := maxAncestorDiff(tree); res != expected {
		t.Log("Input:", tree.String())
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestMaxAncestorDiff(t *testing.T) {
	test(t, []int{
		8,
		3, 10,
		1, 6, null, 14,
		null, null, 4, 7, 13,
	}, 7)
	test(t, []int{
		1,
		null, 2,
		null, null, null, 0,
		null, null, null, null, null, null, 3,
	}, 3)
	test(t, []int{
		8,
		null, 1,
		null, null, 5, 6,
		null, null, null, null, 2, 4, 0, null,
		null, null, null, null, null, null, null, null, 7, 3,
	}, 8)
}
