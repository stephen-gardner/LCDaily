// Problem: https://leetcode.com/problems/sort-characters-by-frequency/
// Results: https://leetcode.com/submissions/detail/853748245/
package main

import "sort"

type Char struct {
	c byte
	n int32
}

func getSortedChars(s string) []Char {
	chars := make([]Char, 1+('z'-'0'))
	for i := range chars {
		chars[i].c = '0' + byte(i)
	}
	for _, c := range s {
		chars[c-'0'].n++
	}
	sort.Slice(chars, func(i, j int) bool {
		return chars[i].n > chars[j].n
	})
	return chars
}

func frequencySort(s string) string {
	chars := getSortedChars(s)
	res := make([]byte, len(s))
	writeIdx := 0
	for _, char := range chars {
		for char.n > 0 {
			res[writeIdx] = char.c
			writeIdx++
			char.n--
		}
	}
	return string(res)
}
