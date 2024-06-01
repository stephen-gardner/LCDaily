// Problem: https://leetcode.com/problems/word-pattern/
// Results: https://leetcode.com/problems/word-pattern/submissions/868851954/
package main

import "strings"

func wordPattern(pattern string, s string) bool {
	bank := map[string]byte{}
	rel := map[byte]string{}
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	for i := range pattern {
		p1, w1 := pattern[i], words[i]
		w2, p2 := rel[p1], bank[w1]
		if (w2 != "" && w1 != w2) || (p2 != 0 && p1 != p2) {
			return false
		}
		rel[p1] = w1
		bank[w1] = p1
	}
	return true
}
