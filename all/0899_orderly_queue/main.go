// Problem: https://leetcode.com/problems/orderly-queue/
// Results: https://leetcode.com/submissions/detail/837851058/
package main

import "sort"

// Booth's Algorithm: https://en.wikipedia.org/wiki/Lexicographically_minimal_string_rotation#Booth's_Algorithm
func leastRotation(s string) int {
	n := len(s)
	f := make([]int, 2*n)
	for i := range f {
		f[i] = -1
	}
	k := 0
	for j := 1; j < 2*n; j++ {
		i := f[j-k-1]
		for i != -1 && s[j%n] != s[(k+i+1)%n] {
			if s[j%n] < s[(k+i+1)%n] {
				k = j - i - 1
			}
			i = f[i]
		}
		if i == -1 && s[j%n] != s[(k+i+1)%n] {
			if s[j%n] < s[(k+i+1)%n] {
				k = j
			}
			f[j-k] = -1
		} else {
			f[j-k] = i + 1
		}
	}
	return k
}

func orderlyQueue(s string, k int) string {
	// If k is greater than 1, then s can be fully sorted
	if k > 1 {
		res := []byte(s)
		sort.Slice(res, func(i, j int) bool {
			return res[i] < res[j]
		})
		return string(res)
	}
	// All we can do is rotate the string
	if start := leastRotation(s); start != 0 {
		s = s[start:] + s[:start]
	}
	return s
}
