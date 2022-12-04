// Problem: https://leetcode.com/problems/determine-if-string-halves-are-alike/
// Results: https://leetcode.com/submissions/detail/852611454/
package main

func halvesAreAlike(s string) bool {
	vowels := ['z' + 1]byte{
		'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1,
		'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1,
	}
	delta := 0
	s1, s2 := s[:len(s)/2], s[len(s)/2:]
	for i := range s1 {
		delta += int(vowels[s1[i]])
		delta -= int(vowels[s2[i]])
	}
	return delta == 0
}
