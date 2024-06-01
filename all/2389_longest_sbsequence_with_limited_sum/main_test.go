package main

import "testing"

func test(t *testing.T, nums, queries, expected []int) {
	res := answerQueries(nums, queries)
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
		t.Log("   Input:", nums)
		t.Log(" Queries:", queries)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestAnswerQueries(t *testing.T) {
	test(t,
		[]int{4, 5, 2, 1},
		[]int{3, 10, 21},
		[]int{2, 3, 4})
	test(t,
		[]int{2, 3, 4, 5},
		[]int{1},
		[]int{0})
	test(t,
		[]int{736411, 184882, 914641, 37925, 214915},
		[]int{331244, 273144, 118983, 118252, 305688, 718089, 665450},
		[]int{2, 2, 1, 1, 2, 3, 3})
}
