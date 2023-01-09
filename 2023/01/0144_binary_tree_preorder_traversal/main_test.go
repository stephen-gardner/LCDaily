package main

import (
	"math"
	"testing"

	"github.com/stephen-gardner/LCDaily/test/tree"
)

const null = math.MaxInt

func test(t *testing.T, nums, expected []int) {
	input := tree.Create(nums, 0)
	res := preorderTraversal(input)
	fail := len(res) != len(expected)
	if !fail {
		for i := range expected {
			if res[i] != expected[i] {
				fail = true
				break
			}
		}
	}
	if fail {
		t.Log("   Input:", input.String())
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestPreorderTraversal(t *testing.T) {
	test(t, []int{1, null, 2, null, null, 3}, []int{1, 2, 3})
	test(t, []int{}, []int{})
	test(t, []int{1}, []int{1})
	test(t, []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 4, 5, 3, 6, 7})
}
