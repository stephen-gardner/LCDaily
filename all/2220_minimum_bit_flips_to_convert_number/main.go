// Problem: https://leetcode.com/problems/minimum-bit-flips-to-convert-number/
// Results: https://leetcode.com/problems/minimum-bit-flips-to-convert-number/submissions/1387916083
package main

import "math/bits"

// Time: O(1)
// Space: O(1)
func minBitFlips(start int, goal int) int {
	return bits.OnesCount(uint(start^goal))
}
