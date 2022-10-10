// Problem: https://leetcode.com/problems/break-a-palindrome/
// Results: https://leetcode.com/submissions/detail/819053616/
package main

func breakPalindrome(str string) string {
	if len(str) == 1 {
		return ""
	}
	res := []byte(str)
	half := len(res) / 2
	for i := 0; i < half; i++ {
		if res[i] != 'a' {
			res[i] = 'a'
			return string(res)
		}
	}
	res[len(res)-1] = 'b'
	return string(res)
}
