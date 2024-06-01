// Problem: https://leetcode.com/problems/daily-temperatures/
// Results: https://leetcode.com/problems/daily-temperatures/submissions/861838453/
package main

func dailyTemperatures(temps []int) []int {
	res := make([]int, len(temps))
	for i := len(temps) - 2; i >= 0; i-- {
		t1 := temps[i]
		for j := i + 1; j < len(temps); j++ {
			t2 := temps[j]
			switch {
			case t1 > t2:
				continue
			case t1 < t2:
				res[i] = j - i
			default:
				if res[j] == 0 {
					res[i] = 0
				} else {
					res[i] = (j - i) + res[j]
				}
			}
			break
		}
	}
	return res
}
