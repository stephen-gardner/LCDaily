// Problem: https://leetcode.com/problems/bag-of-tokens/
// Results: https://leetcode.com/submissions/detail/797560605/
package main

import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	score, bestScore := 0, 0
	for len(tokens) > 0 {
		if power >= tokens[0] {
			score++
			power -= tokens[0]
			tokens = tokens[1:]
		} else if score > 0 {
			score--
			power += tokens[len(tokens)-1]
			tokens = tokens[:len(tokens)-1]
		} else {
			break
		}
		if score > bestScore {
			bestScore = score
		}
	}
	return bestScore
}
