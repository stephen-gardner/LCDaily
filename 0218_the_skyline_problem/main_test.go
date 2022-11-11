package main

import "testing"

func test(t *testing.T, buildings, expected [][]int) {
	dupe := make([][]int, len(buildings))
	for i := range buildings {
		dupe[i] = make([]int, len(buildings[i]))
		for j := range buildings[i] {
			dupe[i][j] = buildings[i][j]
		}
	}
	res := getSkyline(dupe)
	fail := len(res) != len(expected)
	if !fail {
		for i := range expected {
			if res[i][0] != expected[i][0] || res[i][1] != expected[i][1] {
				fail = true
				break
			}
		}
	}
	if fail {
		t.Log("Buildings:", buildings)
		t.Log("   Result:", res)
		t.Log(" Expected:", expected)
		t.FailNow()
	}
}

func TestGetSkyline(t *testing.T) {
	test(
		t,
		[][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}},
		[][]int{{2, 10}, {3, 15}, {7, 12}, {12, 0}, {15, 10}, {20, 8}, {24, 0}},
	)
	test(
		t,
		[][]int{{0, 2, 3}, {2, 5, 3}},
		[][]int{{0, 3}, {5, 0}},
	)
	test(
		t,
		[][]int{{1, 2, 1}, {1, 2, 2}, {1, 2, 3}},
		[][]int{{1, 3}, {2, 0}},
	)
	test(
		t,
		[][]int{{1, 20, 1}, {1, 21, 2}, {1, 22, 3}},
		[][]int{{1, 3}, {22, 0}},
	)
	test(
		t,
		[][]int{{1, 2, 1}, {1, 2, 2}, {1, 2, 3}, {2, 3, 1}, {2, 3, 2}, {2, 3, 3}},
		[][]int{{1, 3}, {3, 0}},
	)
}
