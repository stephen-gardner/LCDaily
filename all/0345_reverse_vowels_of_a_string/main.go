// Problem: https://leetcode.com/problems/reverse-vowels-of-a-string/
// Results: https://leetcode.com/submissions/detail/836469050/
package main

var vowels = [128]bool{
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

func reverseVowels(s string) string {
	res := []byte(s)
	start, end := 0, len(res)-1
	for start < end {
		for start < end && !vowels[res[start]] {
			start++
		}
		for start < end && !vowels[res[end]] {
			end--
		}
		if start < end {
			res[start], res[end] = res[end], res[start]
			start++
			end--
		}
	}
	return string(res)
}
