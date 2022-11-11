package main

func numRollsToTarget(n int, k int, target int) int {
	if n > target || n*k < target {
		return 0
	}
}
