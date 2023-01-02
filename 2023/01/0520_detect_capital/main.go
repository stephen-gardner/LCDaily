// Problem: https://leetcode.com/problems/detect-capital/
// Results: https://leetcode.com/problems/detect-capital/submissions/869410341/
package main

func detectCapitalUse(word string) bool {
	caps := 0
	for _, c := range word {
		if c < 'a' {
			caps++
		}
	}
	return caps == 0 || caps == len(word) || (caps == 1 && word[0] < 'a')
}
