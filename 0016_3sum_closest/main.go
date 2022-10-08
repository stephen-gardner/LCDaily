// Problem: https://leetcode.com/problems/3sum-closest/
// Results: https://leetcode.com/submissions/detail/817680885/
package main

import (
	"math"
	"sort"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	bestDelta := math.MaxInt
	for left := 0; left < len(nums)-2; left++ {
		mid, right := left+1, len(nums)-1
		for mid < right {
			delta := target - (nums[left] + nums[mid] + nums[right])
			if abs(delta) < abs(bestDelta) {
				bestDelta = delta
				if delta == 0 {
					left = len(nums)
					break
				}
			}
			if delta > 0 {
				mid++
			} else {
				right--
			}
		}
	}
	return target - bestDelta
}
