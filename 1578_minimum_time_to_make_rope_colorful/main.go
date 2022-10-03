// Problem: https://leetcode.com/problems/minimum-time-to-make-rope-colorful/
// Results: https://leetcode.com/submissions/detail/814025224/
package main

func minCost(colors string, neededTime []int) int {
	res, shortest := 0, neededTime[0]
	for i := 1; i < len(colors); i++ {
		if colors[i-1] == colors[i] {
			if neededTime[i] < shortest {
				res += neededTime[i]
			} else {
				res += shortest
				shortest = neededTime[i]
			}
		} else {
			shortest = neededTime[i]
		}
	}
	return res
}
