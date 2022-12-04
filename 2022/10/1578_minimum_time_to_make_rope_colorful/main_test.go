package main

import (
	"fmt"
	"strings"
	"testing"
)

func test(t *testing.T, colors string, neededTime []int, expected int) {
	if res := minCost(colors, neededTime); res != expected {
		var sb strings.Builder
		sb.WriteString("[ ")
		for _, c := range colors {
			sb.WriteByte(' ')
			sb.WriteRune(c)
			sb.WriteByte(' ')
		}
		sb.WriteByte(']')
		t.Log("Colors:", sb.String())
		sb.Reset()
		sb.WriteString("[ ")
		for _, n := range neededTime {
			sb.WriteString(fmt.Sprintf("%2d ", n))
		}
		sb.WriteByte(']')
		t.Log("  Time:", sb.String())
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestMinCost(t *testing.T) {
	test(t, "abaac", []int{1, 2, 3, 4, 5}, 3)
	test(t, "abc", []int{1, 2, 3}, 0)
	test(t, "aabaa", []int{1, 2, 3, 4, 1}, 2)
	test(t, "aaabbbabbbb", []int{3, 5, 10, 7, 5, 3, 5, 5, 4, 8, 1}, 26)
}
