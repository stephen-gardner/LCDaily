// Problem: https://leetcode.com/problems/minimum-rounds-to-complete-all-tasks/
// Results: https://leetcode.com/problems/minimum-rounds-to-complete-all-tasks/submissions/870911174/
package main

func minimumRounds(tasks []int) int {
	pool := map[int]int{}
	for i := 0; i < len(tasks); i++ {
		pool[tasks[i]]++
	}
	rounds := 0
	for _, n := range pool {
		if n == 1 {
			return -1
		} else if n%3 == 0 {
			rounds += n / 3
		} else {
			rounds += 1 + (n / 3)
		}
	}
	return rounds
}
