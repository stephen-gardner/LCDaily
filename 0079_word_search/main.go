// Problem: https://leetcode.com/problems/word-search/
// Results: https://leetcode.com/submissions/detail/848892425/
package main

var dirs = [4][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func isPossible(board [][]byte, word string, m, n byte) bool {
	counts := map[byte]int{}
	for y := byte(0); y < m; y++ {
		for x := byte(0); x < n; x++ {
			counts[board[y][x]]++
		}
	}
	for i := range word {
		counts[word[i]]--
	}
	for _, n := range counts {
		if n < 0 {
			return false
		}
	}
	return true
}

func exist(board [][]byte, word string) bool {
	m, n, end := byte(len(board)), byte(len(board[0])), byte(len(word)-1)
	if !isPossible(board, word, m, n) {
		return false
	}
	var hasWord func(byte, byte, byte) bool
	hasWord = func(oy, ox, i byte) bool {
		if i == end {
			return true
		}
		board[oy][ox] = 0
		i++
		for _, dir := range dirs {
			y, x := oy+byte(dir[0]), ox+byte(dir[1])
			if y < m && x < n && board[y][x] == word[i] && hasWord(y, x, i) {
				return true
			}
		}
		board[oy][ox] = word[i-1]
		return false
	}
	for y := byte(0); y < m; y++ {
		for x := byte(0); x < n; x++ {
			if board[y][x] == word[0] && hasWord(y, x, 0) {
				return true
			}
		}
	}
	return false
}
