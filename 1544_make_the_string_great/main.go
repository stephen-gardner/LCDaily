// Problem: https://leetcode.com/problems/make-the-string-great/
// Results: https://leetcode.com/submissions/detail/839101369/
package main

func makeGood(s string) string {
	res := make([]byte, len(s))
	writeIdx := 0
	for i := 0; i < len(s); i++ {
		res[writeIdx] = s[i]
		if writeIdx > 0 && res[writeIdx] == (res[writeIdx-1]^' ') {
			writeIdx--
		} else {
			writeIdx++
		}
	}
	return string(res[:writeIdx])
}
