package main

import (
	"slices"
	"testing"

	"github.com/stephen-gardner/LCDaily/test"
)

func TestSortColors(t *testing.T) {
	test := func(nums []int) {
		res := test.CloneSlice(nums)
		sortColors(res)
		expected := test.CloneSlice(nums)
		slices.Sort(expected)
		for i := range expected {
			if res[i] != expected[i] {
				t.Log("Nums:", nums)
				t.Log("Result:", res)
				t.Log("Expected:", expected)
				t.FailNow()
			}
		}
	}
	test([]int{2, 0, 2, 1, 1, 0})
	test([]int{2, 0, 1})
}
