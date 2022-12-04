// Problem: https://leetcode.com/problems/numbers-with-same-consecutive-differences/
// Results: https://leetcode.com/submissions/detail/790130241/
package main

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func constructNumbers(res *[]int, digits [10][]int, curr, numDigits, maxDigits int) {
	if numDigits == maxDigits {
		*res = append(*res, curr)
		return
	}
	for _, nextDigit := range digits[curr%10] {
		constructNumbers(res, digits, (curr*10)+nextDigit, numDigits+1, maxDigits)
	}
}

func numsSameConsecDiff(n int, k int) []int {
	res := []int{}
	digits := [10][]int{}
	for first := 0; first < 10; first++ {
		for second := 0; second < 10; second++ {
			if abs(first-second) == k {
				digits[first] = append(digits[first], second)
			}
		}
	}
	for i := 1; i < 10; i++ {
		constructNumbers(&res, digits, i, 1, n)
	}
	return res
}
