// Problem: https://leetcode.com/problems/replace-words/
// Results: https://leetcode.com/problems/replace-words/submissions/1280714255
package main

import "strings"

type Trie struct {
	index [26]*Trie
	word  bool
}

func (trie *Trie) Insert(word string) {
	for i := range word {
		c := word[i] - 'a'
		if trie.index[c] == nil {
			trie.index[c] = &Trie{}
		}
		trie = trie.index[c]
	}
	trie.word = true
}

func (trie *Trie) FindShortest(sentence string, pos int) (string, int) {
	var word string
	i := pos
	for i < len(sentence) {
		c := sentence[i] - 'a'
		if c >= 26 || trie.index[c] == nil {
			break
		}
		trie = trie.index[c]
		if trie.word {
			word = sentence[pos : i+1]
			break
		}
		i++
	}
	for i < len(sentence) && sentence[i] != ' ' {
		i++
	}
	if word == "" {
		word = sentence[pos:i]
	}
	return word, i
}

// Time: O(n*m)
// Space: O(n*m)
//
//	n = length of dictionary
//	m = average length of entries
func replaceWords(dictionary []string, sentence string) string {
	var sb strings.Builder
	trie := &Trie{}
	for _, word := range dictionary {
		trie.Insert(word)
	}
	for i := 0; i < len(sentence); i++ {
		var word string
		word, i = trie.FindShortest(sentence, i)
		sb.WriteString(word)
		if i+1 < len(sentence) {
			sb.WriteByte(' ')
		}
	}
	return sb.String()
}
