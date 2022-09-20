// Problem: https://leetcode.com/problems/trapping-rain-water/
// Results: https://leetcode.com/submissions/detail/804761796/
package main

func trap(heights []int) int {
	waterLevel, totalVolume := 0, 0
	left, right := 0, len(heights)-1
	maxLeft, maxRight := heights[left], heights[right]
	for left < right {
		// As we need the minimum height to calculate the water level, we can
		//	use the highest height previously seen for a given side by
		//	processing whichever side is shorter
		if maxLeft < maxRight {
			left++
			waterLevel = maxLeft - heights[left]
			if heights[left] > maxLeft {
				maxLeft = heights[left]
			}
		} else {
			right--
			waterLevel = maxRight - heights[right]
			if heights[right] > maxRight {
				maxRight = heights[right]
			}
		}
		if waterLevel > 0 {
			totalVolume += waterLevel
		}
	}
	return totalVolume
}
