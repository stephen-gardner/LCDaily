// Problem: https://leetcode.com/problems/crawler-log-folder/
// Results: https://leetcode.com/problems/crawler-log-folder/submissions/1316045213
package main

// Time: O(n)
// Space: O(1)
func minOperations(logs []string) int {
	depth := 0
	for _, op := range logs {
		switch op {
		case "./":
			continue
		case "../":
			if depth > 0 {
				depth--
			}
		default:
			depth++
		}
	}
	return depth
}
