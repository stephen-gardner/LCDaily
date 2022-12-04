// Problen: https://leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words/
// Results: https://leetcode.com/submissions/detail/835866293/
package main

func longestPalindrome(words []string) int {
	length := 0
	seen := [26][26]int{}
	for _, w := range words {
		// Both (XX + XX) and (XY + YX) words can pad the ends
		a, b := w[0]-'a', w[1]-'a'
		if seen[b][a] > 0 {
			length += 4
			seen[b][a]--
		} else {
			seen[a][b]++
		}
	}
	// If an XX word has an odd amount of occurrences, it may serve as the
	//	center piece, but there's only room for one
	for i := 0; i < 26; i++ {
		if seen[i][i] > 0 {
			length += 2
			break
		}
	}
	return length
}
