// Problem: https://leetcode.com/problems/find-common-characters/
// Results: https://leetcode.com/problems/find-common-characters/submissions/1278687963
package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func countChars(counts []int, w string) {
	for _, c := range w {
		counts[int(c)-'a']++
	}
}

// Time: O(n)
// Space: O(1)
func commonChars(words []string) []string {
	countsA, countsB := make([]int, 26), make([]int, 26)
	countChars(countsA, words[0])
	for i := 1; i < len(words); i++ {
		countChars(countsB, words[i])
		for j := range countsA {
			countsA[j] = min(countsA[j], countsB[j])
			countsB[j] = 0
		}
	}
	res := []string{}
	for i, n := range countsA {
		c := rune(i) + 'a'
		for j := 0; j < n; j++ {
			res = append(res, string(c))
		}
	}
	return res
}
