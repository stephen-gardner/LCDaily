// Problem: https://leetcode.com/problems/the-number-of-weak-characters-in-the-game/
// Results: https://leetcode.com/submissions/detail/795217587/
package main

import "sort"

const (
	Attack = iota
	Defense
)

func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][Attack] == properties[j][Attack] {
			return properties[i][Defense] < properties[j][Defense]
		}
		return properties[i][Attack] > properties[j][Attack]
	})
	weak := 0
	highDef := 0
	for _, player := range properties {
		if player[Defense] < highDef {
			weak++
		} else if player[Defense] > highDef {
			highDef = player[Defense]
		}
	}
	return weak
}
