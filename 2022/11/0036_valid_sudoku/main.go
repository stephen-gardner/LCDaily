// Problem: https://leetcode.com/problems/valid-sudoku/
// Results: https://leetcode.com/submissions/detail/848345362/
package main

func isValidSudoku(board [][]byte) bool {
	var row, col, box [9]int
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			c := int(board[y][x])
			if c == '.' {
				continue
			}
			c = 1 << (c - '1')
			boxNum := ((y / 3) * 3) + x/3
			if row[y]&c != 0 || col[x]&c != 0 || box[boxNum]&c != 0 {
				return false
			}
			row[y] |= c
			col[x] |= c
			box[boxNum] |= c
		}
	}
	return true
}
