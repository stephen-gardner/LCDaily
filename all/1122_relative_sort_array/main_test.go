package main

import "testing"

func dupe(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	return res
}

func TestRelativeSortArray(t *testing.T) {
	test := func(arr1, arr2, expected []int) {
		res := relativeSortArray(dupe(arr1), dupe(arr2))
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
			t.Log("arr1:", arr1)
			t.Log("arr2:", arr2)
			t.Log("Results:", res)
			t.Log("Expected:", expected)
			t.FailNow()
		}
	}
	test(
		[]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
		[]int{2, 1, 4, 3, 9, 6},
		[]int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
	)
	test(
		[]int{28, 6, 22, 8, 44, 17},
		[]int{22, 28, 8, 6},
		[]int{22, 28, 8, 6, 17, 44},
	)
}
