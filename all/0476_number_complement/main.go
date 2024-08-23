// Problem: https://leetcode.com/problems/number-complement/
// Results: https://leetcode.com/problems/number-complement/submissions/1365182819
package main

import "math/bits"

// Time: O(1)
// Space: O(1)
// Example:
//
//		 0000 0101 (5)
//			XOR
//		 0000 0111 ((1 << 3) - 1)
//	         =
//		 0000 0010 (2)
func findComplement(num int) int {
	return ((1 << bits.Len(uint(num))) - 1) ^ num
}
