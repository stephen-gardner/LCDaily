// Problem: https://leetcode.com/problems/word-search-ii/
// Results: https://leetcode.com/submissions/detail/837114066/
package main

type Trie struct {
	index [26]*Trie
	word  string
}

func (trie *Trie) Insert(word string) {
	for i := range word {
		c := int(word[i]) - 'a'
		if trie.index[c] == nil {
			trie.index[c] = &Trie{}
		}
		trie = trie.index[c]
	}
	trie.word = word
}

func (trie *Trie) IsEmpty() bool {
	if len(trie.word) > 0 {
		return false
	}
	for i := range trie.index {
		if trie.index[i] != nil {
			return false
		}
	}
	return true
}

func (trie *Trie) Remove(word string) bool {
	if len(word) == 0 {
		trie.word = ""
		return trie.IsEmpty()
	}
	next := trie.index[word[0]-'a']
	if next.Remove(word[1:]) {
		trie.index[word[0]-'a'] = nil
		return trie.IsEmpty()
	}
	return false
}

func explore(res *[]string, root, curr *Trie, board [][]byte, y, x int) {
	c := int(board[y][x])
	if c == 0 || curr.index[c-'a'] == nil {
		return
	}
	curr = curr.index[c-'a']
	if len(curr.word) > 0 {
		*res = append(*res, curr.word)
		root.Remove(curr.word)
	}
	board[y][x] = 0
	if y+1 < len(board) {
		explore(res, root, curr, board, y+1, x)
	}
	if y-1 >= 0 {
		explore(res, root, curr, board, y-1, x)
	}
	if x+1 < len(board[0]) {
		explore(res, root, curr, board, y, x+1)
	}
	if x-1 >= 0 {
		explore(res, root, curr, board, y, x-1)
	}
	board[y][x] = byte(c)
}

func findWords(board [][]byte, words []string) []string {
	trie := &Trie{}
	for _, w := range words {
		trie.Insert(w)
	}
	res := []string{}
	for y := range board {
		for x := range board[0] {
			explore(&res, trie, trie, board, y, x)
		}
	}
	return res
}
