// Problem: https://leetcode.com/problems/find-duplicate-file-in-system/
// Results: https://leetcode.com/submissions/detail/803690508/
package main

import "strings"

func findDuplicate(paths []string) [][]string {
	files := map[string][]string{}
	var sb strings.Builder
	for _, data := range paths {
		i := 0
		for data[i] != ' ' {
			i++
		}
		dirEnd := i
		i++
		for j := i; i < len(data); i = j + 2 {
			sb.WriteString(data[:dirEnd])
			sb.WriteByte('/')
			for data[j] != '(' {
				j++
			}
			sb.WriteString(data[i:j])
			contentStart := j + 1
			for data[j] != ')' {
				j++
			}
			content := data[contentStart:j]
			files[content] = append(files[content], sb.String())
			sb.Reset()
		}
	}
	res := [][]string{}
	for _, paths := range files {
		if len(paths) > 1 {
			res = append(res, paths)
		}
	}
	return res
}
