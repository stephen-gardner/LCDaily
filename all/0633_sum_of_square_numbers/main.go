// Problem: https://leetcode.com/problems/sum-of-square-numbers/
// Results: https://leetcode.com/problems/sum-of-square-numbers/submissions/1291596490
package main

// Time: O(Sqrt(c))
// Space: O(Sqrt(c))
// https://en.wikipedia.org/wiki/Sum_of_two_squares_theorem
func judgeSquareSum(c int) bool {
	primes := map[int]int{}
	i := 2
	for i*i <= c {
		if c%i == 0 {
			c /= i
			primes[i]++
			continue
		}
		i++
	}
	primes[c]++
	for p, n := range primes {
		if n&1 == 1 && p%4 == 3 {
			return false
		}
	}
	return true
}
