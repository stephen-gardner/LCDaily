package main

import (
	"math"
	"testing"

	"github.com/stephen-gardner/LCDaily/test/tree"
)

const null = math.MaxInt

func test(t *testing.T, nums []int, expected int) {
	dupe := tree.Create(nums, 0)
	tree := tree.Create(nums, 0)
	if res := maxProduct(dupe); res != expected {
		t.Log("   Input:", tree.String())
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestMaxProduct(t *testing.T) {
	test(t, []int{
		1,
		2, 3,
		4, 5, 6,
	}, 110)
	test(t, []int{
		1,
		null, 2,
		null, null, 3, 4,
		null, null, null, null, null, null, 5, 6,
	}, 90)
}
