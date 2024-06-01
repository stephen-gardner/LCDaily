// Problem: https://leetcode.com/problems/find-k-closest-elements/
// Results: https://leetcode.com/submissions/detail/811192610/
package main

func diff(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}

func findClosestElements(arr []int, k int, x int) []int {
	if len(arr) == k {
		return arr
	}
	bestStart, bestScore := 0, 0
	for i := 0; i < k; i++ {
		bestScore += diff(arr[i], x)
	}
	for i, score := 1, bestScore; i+k <= len(arr); i++ {
		score = score - diff(arr[i-1], x) + diff(arr[i+(k-1)], x)
		if score < bestScore {
			bestStart, bestScore = i, score
		} else if score > bestScore {
			break
		}
	}
	return arr[bestStart : bestStart+k]
}
