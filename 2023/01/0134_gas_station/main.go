// Problem: https://leetcode.com/problems/gas-station/
// Results: https://leetcode.com/problems/gas-station/submissions/873237155/
package main

func canCompleteCircuit(gas []int, cost []int) int {
	startIdx := 0
	total := 0
	for i, curr := 0, 0; i < len(gas); i++ {
		net := gas[i] - cost[i]
		total += net
		curr += net
		// If current sum is < 0, then the route leading up to it can't
		//	possibly be the start
		if curr < 0 {
			startIdx = i + 1
			curr = 0
		}
	}
	if total >= 0 && startIdx < len(gas) {
		return startIdx
	}
	return -1
}
