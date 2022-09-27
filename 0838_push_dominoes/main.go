// Problem: https://leetcode.com/problems/push-dominoes/
// Result: https://leetcode.com/submissions/detail/809653181/
package main

func pushDominoes(dominoes string) string {
	res := []byte(dominoes)
	for i := range res {
		if res[i] != '.' {
			continue
		}
		// Find span of standing dominoes
		start, end := i, i
		for end+1 < len(res) && res[end+1] == '.' {
			end++
		}
		next := end
		// Find leaning
		leanLeft, leanRight := false, false
		if start-1 >= 0 && res[start-1] == 'R' {
			leanRight = true
		}
		if end+1 < len(res) && res[end+1] == 'L' {
			leanLeft = true
		}
		// Work both ends towards middle
		for start < end && (leanRight || leanLeft) {
			if leanRight {
				res[start] = 'R'
				start++
			}
			if leanLeft {
				res[end] = 'L'
				end--
			}
		}
		// Middle piece
		if leanLeft && !leanRight {
			res[start] = 'L'
		} else if !leanLeft && leanRight {
			res[start] = 'R'
		}
		i = next
	}
	return string(res)
}
