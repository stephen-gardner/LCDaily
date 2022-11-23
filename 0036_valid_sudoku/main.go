// Problem: https://leetcode.com/problems/valid-sudoku/
// Results: https://leetcode.com/submissions/detail/848336369/
package main

func isValidSudoku(board [][]byte) bool {
	var row, col, box [9][9]bool
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			c := board[y][x]
			if c == '.' {
				continue
			}
			c -= '1'
			boxNum := ((y / 3) * 3) + x/3
			if row[y][c] || col[x][c] || box[boxNum][c] {
				return false
			}
			row[y][c] = true
			col[x][c] = true
			box[boxNum][c] = true
		}
	}
	return true
}
