package main

import "testing"

func test(t *testing.T, plantTime, growTime []int, expected int) {
	if res := earliestFullBloom(plantTime, growTime); res != expected {
		t.Log("Plant Time:", plantTime)
		t.Log(" Grow Time:", growTime)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestEarliestFullBloom(t *testing.T) {
	test(t, []int{1, 4, 3}, []int{2, 3, 1}, 9)
	test(t, []int{1, 2, 3, 2}, []int{2, 1, 2, 1}, 9)
	test(t, []int{1}, []int{1}, 2)
	test(
		t,
		[]int{15, 29, 24, 8, 14, 26, 12, 15, 27, 2, 5, 10, 7, 17, 9, 5, 9, 21, 11, 13, 13, 2, 1, 17, 11},
		[]int{26, 20, 10, 9, 8, 27, 1, 19, 13, 22, 10, 10, 21, 14, 17, 14, 17, 30, 3, 3, 14, 16, 7, 12, 25},
		324,
	)

}
