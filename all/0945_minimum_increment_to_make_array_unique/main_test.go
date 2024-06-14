package main

import (
	"testing"

	"github.com/stephen-gardner/LCDaily/test"
)

func TestMinIncrementForUnique(t *testing.T) {
	test := func(nums []int, expected int) {
		if res := minIncrementForUnique(test.CloneSlice(nums)); res != expected {
			t.Log("nums:", nums)
			t.Fatalf("%d != %d (expected)", res, expected)
		}
	}
	test([]int{1, 2, 2}, 1)
	test([]int{3, 2, 1, 2, 1, 7}, 6)
}
