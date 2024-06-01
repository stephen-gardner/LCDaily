// Problem: https://leetcode.com/problems/reverse-words-in-a-string/
// Results: https://leetcode.com/submissions/detail/842311789/
package main

func reverse(res []byte, start, end int) {
	for mid := (end - start) / 2; mid >= 0; mid-- {
		res[start+mid], res[end-mid] = res[end-mid], res[start+mid]
	}
}

func reverseWords(s string) string {
	res := []byte(s)
	// Remove extra spaces
	i, writeIdx := 0, 0
	for i < len(res) {
		// Skip all spaces
		for i < len(res) && res[i] == ' ' {
			i++
		}
		if i == len(res) {
			break
		}
		// Add single space between words
		if writeIdx > 0 {
			res[writeIdx] = ' '
			writeIdx++
		}
		for i < len(res) && res[i] != ' ' {
			res[writeIdx] = res[i]
			writeIdx++
			i++
		}
	}
	res = res[:writeIdx]
	i = 0
	// Reverse individual words
	for i < len(res) {
		end := i
		for end < len(res) && res[end] != ' ' {
			end++
		}
		reverse(res, i, end-1)
		i = end + 1
	}
	// Reverse entire string
	reverse(res, 0, len(res)-1)
	return string(res)
}
