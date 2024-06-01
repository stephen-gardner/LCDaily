// Problem: https://leetcode.com/problems/check-if-the-sentence-is-pangram/
// Results: https://leetcode.com/submissions/detail/824129718/
package main

func checkIfPangram(sentence string) bool {
	var chars int32
	for _, c := range sentence {
		chars |= 1 << (c - 'a')
		if chars == 0x3FFFFFF {
			// 26 bits set
			return true
		}
	}
	return false
}
