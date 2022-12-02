// Problem: https://leetcode.com/problems/determine-if-two-strings-are-close/
// Results: https://leetcode.com/submissions/detail/853206976/
package main

import "sort"

func closeStrings(w1 string, w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}
	counts1 := make([]int, 26)
	counts2 := make([]int, 26)
	for _, c := range w1 {
		counts1[c-'a']++
	}
	for _, c := range w2 {
		counts2[c-'a']++
	}
	// Check that both words have the same letters
	for i := 0; i < 26; i++ {
		if counts1[i] > 0 || counts2[i] > 0 {
			if counts1[i] == 0 || counts2[i] == 0 {
				return false
			}
		}
	}
	// Check that both words share the same sized sets of letters
	sort.Ints(counts1)
	sort.Ints(counts2)
	for i := 0; i < 26; i++ {
		if counts1[i] != counts2[i] {
			return false
		}
	}
	return true
}
