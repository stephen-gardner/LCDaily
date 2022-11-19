// Problem: https://leetcode.com/problems/erect-the-fence/
// Results: https://leetcode.com/submissions/detail/846308419/
package main

import "sort"

const (
	X = 0
	Y = 1
)

func eraseInterior(trees, lower, upper [][]int) [][]int {
	present := map[int16]bool{}
	for _, t := range lower {
		present[int16((t[0]<<8)|t[1])] = true
	}
	for _, t := range upper {
		present[int16((t[0]<<8)|t[1])] = true
	}
	writeIdx := 0
	for _, t := range trees {
		if present[int16((t[0]<<8)|t[1])] {
			trees[writeIdx] = t
			writeIdx++
		}
	}
	return trees[:writeIdx]
}

// The cross product is used to determine the sign of the acute angle defined
// 	by three points (origin, endA, endB). It corresponds to the direction
// 	(upward or downward) of the cross product of the two coplanar vectors
// 	defined by the two pairs of points (origin, endA) and (origin, endB).
//	The sign of the acute angle is the sign of the cross product of the two
//	vectors.
//
// 		P = (x2 - x1)(y3 - y1) - (y2-y1)(x3 - x1)
// 			+P = counter-clockwise
// 			-P = clockwise
//
// Source: https://en.wikipedia.org/wiki/Cross_product#Computational_geometry
func crossProduct(origin, endA, endB []int) int {
	return ((endA[X] - origin[X]) * (endB[Y] - origin[Y])) - ((endA[Y] - origin[Y]) * (endB[X] - origin[X]))
}

// Monotone chain algorithm O(n log n)
// Source: https://en.wikibooks.org/wiki/Algorithm_Implementation/Geometry/Convex_hull/Monotone_chain
func outerTrees(trees [][]int) [][]int {
	if len(trees) <= 3 {
		return trees
	}
	sort.Slice(trees, func(i, j int) bool {
		if trees[i][X] == trees[j][X] {
			return trees[i][Y] < trees[j][Y]
		}
		return trees[i][X] < trees[j][X]
	})
	lower, upper := [][]int{}, [][]int{}
	for _, t := range trees {
		for len(lower) >= 2 && crossProduct(lower[len(lower)-2], lower[len(lower)-1], t) < 0 {
			lower = lower[:len(lower)-1]
		}
		lower = append(lower, t)
		for len(upper) >= 2 && crossProduct(upper[len(upper)-2], upper[len(upper)-1], t) > 0 {
			upper = upper[:len(upper)-1]
		}
		upper = append(upper, t)
	}
	return eraseInterior(trees, lower, upper)
}
