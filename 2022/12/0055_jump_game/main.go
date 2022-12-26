// Problem: https://leetcode.com/problems/jump-game/
// Results: https://leetcode.com/problems/jump-game/submissions/865543775/
package main

func canJump(nums []int) bool {
	dst := len(nums) - 1
	reach := make([]bool, len(nums))
	reach[0] = true
	for i := 0; i < len(nums); i++ {
		if !reach[i] {
			continue
		}
		limit := i + nums[i]
		if limit >= dst {
			return true
		}
		for j := i; j <= limit; j++ {
			reach[j] = true
		}
	}
	return false
}
