// Problem: https://leetcode.com/problems/sum-of-subarray-minimums/
// Results: https://leetcode.com/submissions/detail/849891644/
package main

const mod = 1e9 + 7

func sumSubarrayMins(arr []int) int {
	res := 0
	stack := []int{-1}
	popLarger := func(idx, n int) {
		sum := 0
		for len(stack) > 1 && n < arr[stack[len(stack)-1]] {
			i := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			sum += (i - stack[len(stack)-1]) * (idx - i) * arr[i]
		}
		res = (res + sum) % mod
	}
	for i, n := range arr {
		popLarger(i, n)
		stack = append(stack, i)
	}
	popLarger(len(arr), 0)
	return res
}
