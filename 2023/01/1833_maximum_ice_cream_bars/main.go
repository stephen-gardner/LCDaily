// Problem: https://leetcode.com/problems/maximum-ice-cream-bars/
// Results: https://leetcode.com/problems/maximum-ice-cream-bars/submissions/872311663/
package main

import "sort"

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	for i, n := range costs {
		if coins -= n; coins < 0 {
			return i
		}
	}
	return len(costs)
}
