// Problem: https://leetcode.com/problems/maximum-69-number/
// Results: https://leetcode.com/submissions/detail/838399033/
package main

func maximum69Number(num int) int {
	power, pos := 1, 0
	for tmp := num; tmp > 0; tmp /= 10 {
		if tmp%10 == 6 {
			pos = power
		}
		power *= 10
	}
	return num - (6 * pos) + (9 * pos)
}
