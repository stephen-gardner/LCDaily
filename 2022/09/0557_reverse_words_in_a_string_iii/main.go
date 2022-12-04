// Problem: https://leetcode.com/problems/reverse-words-in-a-string-iii/
// Results: https://leetcode.com/submissions/detail/805714272/
package main

func reverseWords(s string) string {
	edit := []byte(s)
	i := 0
	for i < len(edit) {
		for i < len(edit) && edit[i] == ' ' {
			i++
		}
		if i == len(edit) {
			break
		}
		end := i + 1
		for end < len(edit) && edit[end] != ' ' {
			end++
		}
		word := edit[i:end]
		half := len(word) / 2
		for j := 0; j < half; j++ {
			word[j], word[len(word)-1-j] = word[len(word)-1-j], word[j]
		}
		i = end
	}
	return string(edit)
}
