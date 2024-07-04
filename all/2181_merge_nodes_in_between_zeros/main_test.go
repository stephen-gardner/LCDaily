package main

import (
	"testing"

	"github.com/stephen-gardner/LCDaily/test/list"
)

func TestMergeNodes(t *testing.T) {
	test := func(input, expected []int) {
		res := list.Array(mergeNodes(list.Create(input)))
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
			t.Log("Input:", input)
			t.Log("Result:", res)
			t.Log("Expected:", expected)
			t.FailNow()
		}
	}
	test([]int{0, 3, 1, 0, 4, 5, 2, 0}, []int{4, 11})
	test([]int{0, 1, 0, 3, 0, 2, 2, 0}, []int{1, 3, 4})
}
