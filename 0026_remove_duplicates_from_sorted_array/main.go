// Problem: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// Results: https://leetcode.com/submissions/detail/841287768/
package main

func removeDuplicates(nums []int) int {
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
