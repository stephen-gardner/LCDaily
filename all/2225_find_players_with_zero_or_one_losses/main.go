// Problem: https://leetcode.com/problems/find-players-with-zero-or-one-losses/
// Results: https://leetcode.com/submissions/detail/851047619/
package main

import (
	"sort"
)

func findWinners(matches [][]int) [][]int {
	status := map[int32]int16{}
	for _, match := range matches {
		status[int32(match[0])] |= 0
		status[int32(match[1])]++
	}
	answer := make([][]int, 2)
	for id, state := range status {
		switch state {
		case 0:
			answer[0] = append(answer[0], int(id))
		case 1:
			answer[1] = append(answer[1], int(id))
		}
	}
	sort.Ints(answer[0])
	sort.Ints(answer[1])
	return answer
}
