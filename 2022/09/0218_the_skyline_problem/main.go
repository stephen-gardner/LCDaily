package main

import (
	"fmt"
	"sort"
)

const (
	Left = iota
	Right
	Height
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sortBuildings(buildings [][]int) {
	sort.Slice(buildings, func(i, j int) bool {
		if buildings[i][Left] == buildings[j][Left] {
			if buildings[i][Right] == buildings[j][Right] {
				return buildings[i][Height] > buildings[j][Height]
			}
			return buildings[i][Right] < buildings[j][Right]
		}
		return buildings[i][Left] < buildings[j][Left]
	})
}

func mergeBuildings(buildings [][]int) [][]int {
	res := [][]int{buildings[0]}
	for i := 1; i < len(buildings); i++ {
		curr, prev := buildings[i], res[len(res)-1]
		if curr[Left] > prev[Right] {
			// No overlap
			if curr[Left]-prev[Right] > 0 {
				// Gap between buildings
				res = append(res, []int{prev[Right], curr[Left], 0})
			}
			res = append(res, curr)
			continue
		}
		if curr[Height] == prev[Height] {
			// Same height, so extend length of previous building
			prev[Right] = max(prev[Right], curr[Right])
		} else if curr[Height] < prev[Height] {
			// Shorter building isn't eclipsed by previous one, so it starts
			//	where the other ends
			if curr[Right] > prev[Right] {
				curr[Left] = prev[Right]
				res = append(res, curr)
			}
		} else {
			if curr[Left] == prev[Left] {
				prev[Height] = curr[Height]
				prev[Right] = max(prev[Right], curr[Right])
			} else {
				res = append(res, curr)
			}
		}
	}
	return res
}

func getSkyline(buildings [][]int) [][]int {
	sortBuildings(buildings)
	fmt.Println(buildings)
	buildings = mergeBuildings(buildings)
	fmt.Println(buildings)
	end := buildings[len(buildings)-1][Right]
	for i, curr := range buildings {
		curr[Right] = curr[Height]
		buildings[i] = curr[:2]
	}
	buildings = append(buildings, []int{end, 0})
	return buildings
}
