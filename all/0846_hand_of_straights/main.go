// Problem: https://leetcode.com/problems/hand-of-straights/
// Results: https://leetcode.com/problems/hand-of-straights/submissions/1279766324
package main

// Time: O(n): It looks crazy with the nested loop, but it's no different from looping through the hand a second time
// Space: O(n)
func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}
	counts := map[int]int{}
	for _, n := range hand {
		counts[n]++
	}
	for _, n := range hand {
		if counts[n] == 0 {
			continue
		}
		for counts[n-1] > 0 {
			n--
		}
		for j := 0; j < groupSize; j++ {
			counts[n+j]--
			if counts[n+j] < 0 {
				return false
			}
		}
	}
	return true
}
