// Problem: https://leetcode.com/problems/perfect-squares/
// Results: https://leetcode.com/submissions/detail/847911185/
package main

import (
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isSquare(n int) bool {
	sr := int(math.Sqrt(float64(n)))
	return sr*sr == n
}

func getSquares(n int) []int {
	res := []int{}
	for i := 1; i < n; i++ {
		if isSquare(i) {
			res = append(res, i)
		}
	}
	res = append(res, n)
	return res
}

func numSquares(n int) int {
	sr := int(math.Sqrt(float64(n)))
	high := sr * sr
	if high == n {
		return 1
	}
	squares := getSquares(high)
	counts := make([]int, n+1)
	for i := 1; i < len(counts); i++ {
		counts[i] = n + 1
	}
	for currAmount := 1; currAmount <= n; currAmount++ {
		for _, sq := range squares {
			if sq > currAmount {
				break
			}
			counts[currAmount] = min(counts[currAmount], 1+counts[currAmount-sq])
		}
	}
	return counts[n]
}
