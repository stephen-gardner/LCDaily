// Problem: https://leetcode.com/problems/earliest-possible-day-of-full-bloom/
// Results: https://leetcode.com/submissions/detail/832418498/
package main

import (
	"sort"
)

type Plant struct {
	plantTime int
	growTime  int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func earliestFullBloom(plantTime []int, growTime []int) int {
	plants := make([]int, len(plantTime))
	for i := range plants {
		plants[i] = i
	}
	sort.Slice(plants, func(i, j int) bool {
		return growTime[plants[i]] > growTime[plants[j]]
	})
	days, grow := plantTime[plants[0]], growTime[plants[0]]
	for i := 1; i < len(plants); i++ {
		p := plants[i]
		days += plantTime[p]
		grow -= plantTime[p]
		grow = max(grow, growTime[p])
	}
	days += grow
	return days
}
