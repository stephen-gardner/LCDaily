// Problem: https://leetcode.com/problems/append-characters-to-string-to-make-subsequence/
// Results: https://leetcode.com/problems/append-characters-to-string-to-make-subsequence/submissions/1276533650
package main

// Time: O(n)
// Space: O(1)
func appendCharacters(s string, t string) int {
	found := 0
	for i := 0; i < len(s) && found < len(t); i++ {
		if s[i] == t[found] {
			found++
		}
	}
	return len(t) - found
}
