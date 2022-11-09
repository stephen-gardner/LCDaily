package main

import "testing"

func test(t *testing.T, nums, expected []int) {
	stockSpanner := Constructor()
	res := make([]int, len(nums))
	for i, price := range nums {
		res[i] = stockSpanner.Next(price)
	}
	for i := range expected {
		if res[i] != expected[i] {
			t.Log("   Input:", nums)
			t.Log("  Result:", res)
			t.Log("Expected:", expected)
			t.FailNow()
		}
	}
}

func TestStockSpanner(t *testing.T) {
	test(
		t,
		[]int{100, 80, 60, 70, 60, 75, 85},
		[]int{1, 1, 1, 2, 1, 4, 6},
	)
	test(
		t,
		[]int{58, 45, 59, 39, 60},
		[]int{1, 1, 3, 1, 5},
	)
	test(
		t,
		[]int{29, 91, 62, 76, 51},
		[]int{1, 2, 1, 2, 1},
	)
	test(
		t,
		[]int{1007, 1464, 6977, 453, 5739},
		[]int{1, 2, 3, 1, 2},
	)
}
