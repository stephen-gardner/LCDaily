// Problem: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
// Results: https://leetcode.com/submissions/detail/840412018/
package main

func removeDuplicates(s string) string {
	res := make([]byte, len(s))
	writeIdx := 0
	for i := range s {
		if writeIdx > 0 && res[writeIdx-1] == s[i] {
			writeIdx--
		} else {
			res[writeIdx] = s[i]
			writeIdx++
		}
	}
	return string(res[:writeIdx])
}
