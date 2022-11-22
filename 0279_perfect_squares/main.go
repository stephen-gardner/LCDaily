// Problem: https://leetcode.com/problems/perfect-squares/
// Results: https://leetcode.com/submissions/detail/848030010/
package main

import (
	"math"
)

func isSquare(n int) bool {
	sr := int(math.Sqrt(float64(n)))
	return sr*sr == n
}

func numSquares(n int) int {
	// Every natural number can be represented as the sum of four integer squares
	// 	(https://en.wikipedia.org/wiki/Lagrange's_four-square_theorem)

	// A natural number can be represented as the sum of three squares of
	// 	integers if and only if n is not of the form n = 4^a(8b + 7) for
	//	nonnegative integers a and b.
	// 	(https://en.wikipedia.org/wiki/Legendre%27s_three-square_theorem)

	// An integer greater than one can be written as a sum of two squares if
	//	and only if its prime decomposition contains no factor p^k, where prime
	//	p ≡ 3 (mod 4) and k is odd.
	//	(https://en.wikipedia.org/wiki/Sum_of_two_squares_theorem)
	for n%4 == 0 {
		n /= 4
	}
	if n%8 == 7 {
		return 4
	}
	// n could be a square
	if isSquare(n) {
		return 1
	}
	// n could be the sum of two squares
	max := int(math.Sqrt(float64(n))) + 1
	for i := 1; i <= max; i++ {
		if isSquare(n - (i * i)) {
			return 2
		}
	}
	return 3
}
