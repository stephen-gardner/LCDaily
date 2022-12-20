package main

import "testing"

func test(t *testing.T, rooms [][]int, expected bool) {
	if canVisitAllRooms(rooms) != expected {
		t.Log("Input:", rooms)
		t.Fatalf("%v != %v (expeced)", !expected, expected)
	}
}

func TestCanVisitAllRooms(t *testing.T) {
	test(t, [][]int{{1}, {2}, {3}, {}}, true)
	test(t, [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, false)
}
