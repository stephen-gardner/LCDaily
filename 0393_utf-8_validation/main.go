// Problem: https://leetcode.com/problems/utf-8-validation/
// Results: https://leetcode.com/submissions/detail/798420663/
package main

func validUtf8(data []int) bool {
	extra := 0
	for len(data) > 0 {
		b := byte(data[0])
		data = data[1:]
		if extra > 0 {
			if b>>6 != 0b10 {
				return false
			}
			extra--
		} else if b <= 0b01111111 {
			continue
		} else if b <= 0b11011111 {
			extra = 1
		} else if b <= 0b11101111 {
			extra = 2
		} else if b <= 0b11110111 {
			extra = 3
		} else {
			return false
		}
	}
	return extra == 0
}
