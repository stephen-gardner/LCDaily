// Problem: https://leetcode.com/problems/average-waiting-time/
// Results: https://leetcode.com/problems/average-waiting-time/submissions/1315642742
package main

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// Time: O(n)
// Space: O(1)
func averageWaitingTime(customers [][]int) float64 {
	cycle := customers[0][0]
	total := 0
	for i := range customers {
		arrival, time := customers[i][0], customers[i][1]
		wait := (cycle - arrival) + time
		total += wait
		if i+1 < len(customers) {
			// Either behind schedule, or finished before next customer arrived
			cycle = max(cycle+time, customers[i+1][0])
		}
	}
	return float64(total) / float64(len(customers))
}
