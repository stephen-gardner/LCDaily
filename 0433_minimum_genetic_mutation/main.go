// Problem: https://leetcode.com/problems/minimum-genetic-mutation/
// Results: https://leetcode.com/submissions/detail/835319456/
package main

func mutate(min *int, bank []string, curr string, endIdx, mutations int) {
	mutations++
	for i, next := range bank {
		if len(next) == 0 {
			continue
		}
		delta := 0
		for j := range next {
			if next[j] != curr[j] {
				delta++
			}
		}
		if delta != 1 {
			continue
		}
		if i == endIdx {
			if *min == -1 || mutations < *min {
				*min = mutations
			}
			return
		}
		bank[i] = ""
		mutate(min, bank, next, endIdx, mutations)
		bank[i] = next
	}
}

func minMutation(start string, end string, bank []string) int {
	endIdx := -1
	for i, s := range bank {
		if s == end {
			endIdx = i
			break
		}
	}
	if endIdx == -1 {
		return -1
	}
	min := -1
	mutate(&min, bank, start, endIdx, 0)
	return min
}
