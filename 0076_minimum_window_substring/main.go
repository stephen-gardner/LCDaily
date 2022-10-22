// Problem: https://leetcode.com/problems/minimum-window-substring/
// Results: https://leetcode.com/submissions/detail/827686423/
package main

type Window struct {
	start  int
	end    int
	unique int
	counts [128]int16
}

func minWindow(s string, t string) string {
	seek := Window{end: len(s) - 1}
	win := Window{}
	for _, c := range t {
		if seek.counts[c] == 0 {
			seek.unique++
		}
		seek.counts[c]++
	}
	bestStart, bestEnd := -1, -1
	for win.end <= seek.end || win.start < win.end {
		for win.end <= seek.end && win.unique < seek.unique {
			win.counts[s[win.end]]++
			if win.counts[s[win.end]] == seek.counts[s[win.end]] {
				win.unique++
			}
			win.end++
		}
		if win.unique != seek.unique {
			break
		}
		for seek.counts[s[win.start]] == 0 {
			win.start++
		}
		if bestStart == -1 || win.end-win.start < bestEnd-bestStart {
			bestStart, bestEnd = win.start, win.end
		}
		win.counts[s[win.start]]--
		if win.counts[s[win.start]] < seek.counts[s[win.start]] {
			win.unique--
		}
		win.start++
	}
	if bestStart == -1 {
		return ""
	}
	return s[bestStart:bestEnd]
}
