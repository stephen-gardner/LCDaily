// Problem: https://leetcode.com/problems/maximum-distance-in-arrays/
// Results: https://leetcode.com/problems/maximum-distance-in-arrays/submissions/1358423536
package main

type Pick struct {
	val int
	idx int
}

// Time: O(n); n = len(arrays)
// Space: O(1)
func maxDistance(arrays [][]int) int {
	min1 := Pick{arrays[0][0], 0}
	min2 := Pick{arrays[1][0], 1}
	if min2.val < min1.val {
		min1, min2 = min2, min1
	}
	max1 := Pick{arrays[0][len(arrays[0])-1], 0}
	max2 := Pick{arrays[1][len(arrays[1])-1], 1}
	if max2.val > max1.val {
		max1, max2 = max2, max1
	}
	for i := 2; i < len(arrays); i++ {
		if arrays[i][0] < min2.val {
			min2 = Pick{arrays[i][0], i}
			if min2.val < min1.val {
				min1, min2 = min2, min1
			}
		}
		if arrays[i][len(arrays[i])-1] > max2.val {
			max2 = Pick{arrays[i][len(arrays[i])-1], i}
			if max2.val > max1.val {
				max1, max2 = max2, max1
			}
		}
	}
	if min1.idx != max1.idx {
		return max1.val - min1.val
	}
	if max1.val-min2.val > max2.val-min1.val {
		return max1.val - min2.val
	}
	return max2.val - min1.val
}
