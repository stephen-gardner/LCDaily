// Problem: https://leetcode.com/problems/integer-to-roman/
// Results: https://leetcode.com/submissions/detail/826278510/
package main

import "strings"

type RomanNumeral struct {
	value  int
	symbol string
}

var numerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	var sb strings.Builder
	for num > 0 {
		for _, rn := range numerals {
			if num >= rn.value {
				sb.WriteString(rn.symbol)
				num -= rn.value
				break
			}
		}
	}
	return sb.String()
}
