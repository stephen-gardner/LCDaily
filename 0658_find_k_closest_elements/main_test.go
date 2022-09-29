package main

import "testing"

func test(t *testing.T, k, x int, arr []int, expected []int) {
	res := findClosestElements(arr, k, x)
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
		t.Logf("k = %d, x = %d", k, x)
		t.Log("   Input:", arr)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestFindClosestElements(t *testing.T) {
	test(
		t, 4, 3,
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 3, 4},
	)
	test(
		t, 4, -1,
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 3, 4},
	)
	test(
		t, 3, 20,
		[]int{-10, 0, 3, 9, 11, 12, 18, 42},
		[]int{11, 12, 18},
	)
	test(
		t, 1, 9,
		[]int{1, 1, 1, 10, 10, 10},
		[]int{10},
	)
	test(
		t, 1, 1,
		[]int{1, 2},
		[]int{1},
	)
	test(
		t, 1, 14,
		[]int{1},
		[]int{1},
	)
}
