// Problem: https://leetcode.com/problems/basic-calculator/
// Results: https://leetcode.com/submissions/detail/846675562/
package main

func calculate(s string) int {
	// Start with a zero in case the first charater is a unary
	nums, ops := []int32{0}, []byte{}
	popOp := func() {
		if ops[len(ops)-1] == '+' {
			nums[len(nums)-2] += nums[len(nums)-1]
		} else {
			nums[len(nums)-2] -= nums[len(nums)-1]
		}
		nums = nums[:len(nums)-1]
		ops = ops[:len(ops)-1]
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
			continue
		case '+', '-':
			ops = append(ops, s[i])
		case '(':
			ops = append(ops, s[i])
			// Skip all spaces immediately to check for unary
			for i < len(s)-1 && s[i+1] == ' ' {
				i++
			}
			// Rather than introduce multiplication, we'll subtract from 0
			if i < len(s)-1 && s[i+1] == '-' {
				nums = append(nums, 0)
			}
		case ')':
			ops = ops[:len(ops)-1]
			if len(ops) > 0 && ops[len(ops)-1] != '(' {
				popOp()
			}
		default:
			n := int32(0)
			for i < len(s) && uint8(s[i])-'0' < 10 {
				n = (n * 10) + int32(s[i]-'0')
				i++
			}
			i--
			nums = append(nums, n)
			// Perform pending addition/subtraction immediately
			if len(ops) > 0 && ops[len(ops)-1] != '(' {
				popOp()
			}
		}
	}
	return int(nums[len(nums)-1])
}
