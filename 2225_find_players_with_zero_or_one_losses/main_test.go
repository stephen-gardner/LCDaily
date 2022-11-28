package main

import (
	"sort"
	"testing"
)

func test(t *testing.T, matches, expected [][]int) {
	res := findWinners(matches)
	fail := len(res[0]) != len(expected[0]) || len(res[1]) != len(expected[1])
	if !fail {
		sort.Ints(res[0])
		sort.Ints(res[1])
		sort.Ints(expected[0])
		sort.Ints(expected[1])
		for i := range expected[0] {
			if res[0][i] != expected[0][i] {
				fail = true
				break
			}
		}
		if !fail {
			for i := range expected[1] {
				if res[1][i] != expected[1][i] {
					fail = true
					break
				}
			}
		}
	}
	if fail {
		t.Log("   Input:", matches)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestFindWinners(t *testing.T) {
	test(t,
		[][]int{
			{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9},
		},
		[][]int{
			{1, 2, 10}, {4, 5, 7, 8},
		})
	test(t,
		[][]int{
			{2, 3}, {1, 3}, {5, 4}, {6, 4},
		},
		[][]int{
			{1, 2, 5, 6}, {},
		})
	test(t,
		[][]int{
			{1, 5}, {2, 5}, {2, 8}, {2, 9}, {3, 8}, {4, 7}, {4, 9}, {5, 7}, {6, 8},
		},
		[][]int{
			{1, 2, 3, 4, 6}, {},
		})
}
