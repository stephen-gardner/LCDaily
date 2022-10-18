package main

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	var sb strings.Builder
	sb.WriteByte('1')
	for i := 1; i < n; i++ {
		curr := sb.String()
		sb.Reset()
		c, count := curr[0], 1
		for j := 1; j < len(curr); j++ {
			if curr[j] == c {
				count++
			} else {
				sb.WriteString(strconv.Itoa(count))
				sb.WriteByte(c)
				c = curr[j]
				count = 1
			}
		}
		sb.WriteString(strconv.Itoa(count))
		sb.WriteByte(c)
	}
	return sb.String()
}
