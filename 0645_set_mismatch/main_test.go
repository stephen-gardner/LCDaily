package main

import "testing"

func test(t *testing.T, nums, expected []int) {
	res := findErrorNums(nums)
	if res[0] != expected[0] || res[1] != expected[1] {
		t.Log("     Nums:", nums)
		t.Log("   Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestFindErrorNums(t *testing.T) {
	test(t, []int{1, 2, 2, 4}, []int{2, 3})
	test(t, []int{1, 1}, []int{1, 2})
	test(t, []int{3, 3, 1}, []int{3, 2})
}
