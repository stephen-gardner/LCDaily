// Problem: https://leetcode.com/problems/count-the-number-of-consistent-strings/
// Results: https://leetcode.com/problems/count-the-number-of-consistent-strings/submissions/1387936901
package main

// Time: O(mn); m = len(words), n = len(words[i])
// Space: O(1)
func countConsistentStrings(allowed string, words []string) int {
	chars := [26]byte{}
	valid := func(word string) bool {
		for _, c := range word {
			if chars[c-'a'] == 0 {
				return false
			}
		}
		return true
	}
	for _, c := range allowed {
		chars[c-'a']++
	}
	count := 0
	for _, word := range words {
		if valid(word) {
			count++
		}
	}
	return count
}
