// Problem: https://leetcode.com/problems/reverse-string/
// Results: https://leetcode.com/problems/reverse-string/submissions/1275593665
package main

// Time:  O(n)
// Space: O(1)
func reverseString(s []byte) {
	mid, end := len(s)/2, len(s)-1
	for i := 0; i < mid; i++ {
		s[i], s[end-i] = s[end-i], s[i]
	}
}
