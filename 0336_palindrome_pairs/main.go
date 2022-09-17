// Problem: https://leetcode.com/problems/palindrome-pairs/
// Results: https://leetcode.com/submissions/detail/802043366/
package main

type Trie struct {
	index       [26]*Trie
	wordIdx     int
	palindromes []int
}

func NewTrie() *Trie {
	return &Trie{wordIdx: -1}
}

func (this *Trie) Insert(word string, idx int) {
	// Insert word in reverse order
	for i := len(word) - 1; i >= 0; i-- {
		c := word[i] - 'a'
		if this.index[c] == nil {
			this.index[c] = NewTrie()
		}
		// If the start of the word [0:i+1] is a palindrome, we store the
		//	occurrence so that it can be effectively skipped
		// e.g., "s" + "lls" will only need to process "s" to find the match
		if isPalindrome(word[:i+1]) {
			this.palindromes = append(this.palindromes, idx)
		}
		this = this.index[c]
	}
	this.wordIdx = idx
}

func (this *Trie) WritePalindromePairs(res *[][]int, word string, wordIdx int) {
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		// Match to words with trailing palindromes
		// e.g., "tabbbbb" will match "bat" (stored in reverse) because the rest
		//	is a palindrome and ignored
		if this.wordIdx != -1 && isPalindrome(word[i:]) {
			*res = append(*res, []int{wordIdx, this.wordIdx})
		}
		if this.index[c] == nil {
			return
		}
		this = this.index[c]
	}
	// Direct match to word in reverse
	// Need to make sure word doesn't match to itself, as it can be a palindrome
	if this.wordIdx != -1 && this.wordIdx != wordIdx {
		*res = append(*res, []int{wordIdx, this.wordIdx})
	}
	// Matches to words with leading palindromes
	for _, idx := range this.palindromes {
		*res = append(*res, []int{wordIdx, idx})
	}
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func palindromePairs(wordArr []string) [][]int {
	words := NewTrie()
	for i, w := range wordArr {
		words.Insert(w, i)
	}
	res := [][]int{}
	for i, w := range wordArr {
		words.WritePalindromePairs(&res, w, i)
	}
	return res
}
