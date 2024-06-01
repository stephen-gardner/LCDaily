// Problem: https://leetcode.com/problems/maximum-bags-with-full-capacity-of-rocks/
// Results: https://leetcode.com/problems/maximum-bags-with-full-capacity-of-rocks/submissions/866076016/
package main

import "sort"

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	space := make([]int, len(capacity))
	for i := range space {
		space[i] = capacity[i] - rocks[i]
	}
	sort.Ints(space)
	full := 0
	for _, n := range space {
		if additionalRocks < n {
			break
		}
		additionalRocks -= n
		full++
	}
	return full
}
