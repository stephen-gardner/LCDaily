package main

import "testing"

func TestXorQueries(t *testing.T) {
	test := func(arr []int, queries [][]int, expected []int) {
		res := xorQueries(arr, queries)
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
			t.Log("arr:", arr)
			t.Log("queries:", queries)
			t.Log("Result:", res)
			t.Log("Expected:", expected)
			t.FailNow()
		}
	}
	test(
		[]int{1, 3, 4, 8},
		[][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}},
		[]int{2, 7, 14, 8},
	)
}
