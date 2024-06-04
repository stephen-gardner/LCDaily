// Problem: https://leetcode.com/problems/longest-palindrome/
// Results: https://leetcode.com/problems/longest-palindrome/submissions/1277650623
package main


// Time: O(n)
// Space: O(1)
func longestPalindrome(s string) int {
	counts := map[rune]int{}
	for _, c := range s {
		counts[c]++
	}
	length := 0
	for _, n := range counts {
		if n&1 == 1 && length&1 == 1 {
			n--
		}
		length += n
	}
	return length
}
