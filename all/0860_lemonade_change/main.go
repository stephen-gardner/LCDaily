// Problem: https://leetcode.com/problems/lemonade-change/
// Results: https://leetcode.com/problems/lemonade-change/submissions/1356054103
package main

// Time: O(n)
// Space: O(1)
func lemonadeChange(bills []int) bool {
	cash := map[int]int{}
	for _, bill := range bills {
		change := bill - 5
		if change > 10 && cash[10] > 0 {
			cash[10]--
			change -= 10
		}
		n := change / 5
		if n > cash[5] {
			return false
		}
		cash[5] -= n
        cash[bill]++
	}
	return true
}