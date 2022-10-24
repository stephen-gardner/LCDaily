// Problem: https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/
// Results: https://leetcode.com/submissions/detail/829136838/
package main

func backtrack(arr []string, counts []int, i int, currLength int, currCount int) int {
	bestLength := currLength
	for i += 1; i < len(arr); i++ {
		if currCount&counts[i] != 0 {
			continue
		}
		length := backtrack(arr, counts, i, currLength+len(arr[i]), currCount|counts[i])
		if length > bestLength {
			bestLength = length
		}
	}
	return bestLength
}

func maxLength(arr []string) int {
	counts := make([]int, len(arr))
	for i, s := range arr {
		for _, c := range s {
			m := 1 << (c - 'a')
			if counts[i]&m != 0 {
				arr[i] = ""
				counts[i] = 0xFFFFFFFF
			} else {
				counts[i] |= m
			}
		}
	}
	return backtrack(arr, counts, -1, 0, 0)
}
