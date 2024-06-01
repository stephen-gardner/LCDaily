package main

import "testing"

func test(t *testing.T, costs []int, coins, expected int) {
	if res := maxIceCream(costs, coins); res != expected {
		t.Log("Costs:", costs)
		t.Log("Coins:", coins)
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestMaxIceCream(t *testing.T) {
	test(t, []int{1, 3, 2, 4, 1}, 7, 4)
	test(t, []int{10, 6, 8, 7, 7, 8}, 5, 0)
	test(t, []int{1, 6, 3, 1, 2, 5}, 20, 6)
}
