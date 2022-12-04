// Problem: https://leetcode.com/problems/toeplitz-matrix/
// Results: https://leetcode.com/submissions/detail/833850368/
package main

func isToeplitzMatrix(matrix [][]int) bool {
	m, n := len(matrix)-1, len(matrix[0])-1
	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			if matrix[y][x] != matrix[y+1][x+1] {
				return false
			}
		}
	}
	return true
}
