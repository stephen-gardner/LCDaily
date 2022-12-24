// Problem: https://leetcode.com/problems/domino-and-tromino-tiling/
// Results: https://leetcode.com/problems/domino-and-tromino-tiling/submissions/864555450/
package main

const mod = int(1e9) + 7

func numTilings(n int) int {
	res := [3]int{1, 2, 5}
	if n <= 3 {
		return res[n-1]
	}
	for i := 3; i < n; i++ {
		ways := ((res[2] * 2) + res[0]) % mod
		res[0], res[1], res[2] = res[1], res[2], ways
	}
	return res[2]
}
