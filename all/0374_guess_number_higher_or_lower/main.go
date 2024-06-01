// Problem: https://leetcode.com/problems/guess-number-higher-or-lower/
// Results: https://leetcode.com/submissions/detail/844265767/
package main

var guess func(int) int

func guessNumber(n int) int {
	low, high := uint32(1), uint32(n)
	for {
		mid := (low + high) / 2
		switch guess(int(mid)) {
		case -1:
			high = mid - 1
		case 0:
			return int(mid)
		case 1:
			low = mid + 1
		}
	}
}
