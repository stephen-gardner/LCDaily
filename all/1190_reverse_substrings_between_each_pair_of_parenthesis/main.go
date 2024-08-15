// Problem: https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/
// Results: https://leetcode.com/problems/reverse-substrings-between-each-pair-of-parentheses/submissions/1319162464
package main

import (
	"bytes"
)

func reverse(data []byte) {
	mid := len(data) / 2
	end := len(data) - 1
	for i := 0; i < mid; i++ {
		data[i], data[end-i] = data[end-i], data[i]
	}
}

func process(data []byte) []byte {
	var buff bytes.Buffer
	for start := 0; start < len(data); start++ {
		switch data[start] {
		case '(':
			end := start
			depth := 1
			for depth > 0 {
				end++
				switch data[end] {
				case '(':
					depth++
				case ')':
					depth--
				}
			}
			sub := process(data[start+1 : end])
			reverse(sub)
			buff.Write(sub)
			start = end
		case ')':
			continue
		default:
			buff.WriteByte(data[start])
		}
	}
	return buff.Bytes()
}

// Time: O(n^2)
// Space: O(n)
func reverseParentheses(s string) string {
	return string(process([]byte(s)))
}
