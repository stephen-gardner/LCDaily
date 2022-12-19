// Problem: https://leetcode.com/problems/evaluate-reverse-polish-notation/
// Results: https://leetcode.com/problems/evaluate-reverse-polish-notation/submissions/861884199/
package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, t := range tokens {
		switch t {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
		default:
			n, _ := strconv.Atoi(t)
			stack = append(stack, n)
			continue
		}
		stack = stack[:len(stack)-1]

	}
	return stack[len(stack)-1]
}
