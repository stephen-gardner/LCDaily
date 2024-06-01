// Problem: https://leetcode.com/problems/score-of-a-string/
// Results: https://leetcode.com/problems/score-of-a-string/submissions/1274544498
package main

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Time:  O(n)
// Space: O(1)
func scoreOfString(s string) int {
	score := 0
	for i := 1; i < len(s); i++ {
		score += abs(int(s[i]) - int(s[i-1]))
	}
	return score
}
