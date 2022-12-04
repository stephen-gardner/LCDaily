package main

import "testing"

func test(t *testing.T, changed, expected []int) {
	res := findOriginalArray(changed)
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
		t.Log("   Input:", changed)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestFindOriginalArray(t *testing.T) {
	test(t, []int{1, 3, 4, 2, 6, 8}, []int{1, 3, 4})
	test(t, []int{6, 3, 0, 1}, []int{})
	test(t, []int{1}, []int{})
	test(t, []int{4, 4, 16, 20, 8, 8, 2, 10}, []int{2, 4, 8, 10})
}
