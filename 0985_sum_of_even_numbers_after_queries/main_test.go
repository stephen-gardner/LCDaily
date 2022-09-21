package main

import "testing"

func test(t *testing.T, nums []int, queries [][]int, expected []int) {
	res := sumEvenAfterQueries(nums, queries)
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
		t.Log("    Nums:", nums)
		t.Log(" Queries:", queries)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestSumEvenAfterQueries(t *testing.T) {
	test(
		t,
		[]int{1, 2, 3, 4},
		[][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}},
		[]int{8, 6, 2, 4},
	)
	test(
		t,
		[]int{1},
		[][]int{{4, 0}},
		[]int{0},
	)
}
