package main

import "testing"

func test(t *testing.T, n, k int, expected []int) {
	res := numsSameConsecDiff(n, k)
	fail := len(res) != len(expected)
	if !fail {
		resMap := map[int]bool{}
		for _, n := range res {
			resMap[n] = true
		}
		for i := 0; i < len(expected); i++ {
			if _, exists := resMap[expected[i]]; !exists {
				fail = true
				break
			}
		}
	}
	if fail {
		t.Log("n =", n, "k =", k)
		t.Log(" Results:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestNumsSameConsecDiff(t *testing.T) {
	test(t, 3, 7, []int{181, 292, 707, 818, 929})
	test(t, 2, 1, []int{10, 12, 21, 23, 32, 34, 43, 45, 54, 56, 65, 67, 76, 78, 87, 89, 98})
}
