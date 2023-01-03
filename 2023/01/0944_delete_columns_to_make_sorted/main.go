// Problem: https://leetcode.com/problems/delete-columns-to-make-sorted/
// Results: https://leetcode.com/problems/delete-columns-to-make-sorted/submissions/870088215/
package main

func minDeletionSize(strs []string) int {
	deletes := 0
	for i := range strs[0] {
		c := byte(0)
		for _, s := range strs {
			if s[i] < c {
				deletes++
				break
			}
			c = s[i]
		}
	}
	return deletes
}
