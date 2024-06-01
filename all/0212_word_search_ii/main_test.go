package main

import "testing"

func test(t *testing.T, board [][]byte, words, expected []string) {
	res := findWords(board, words)
	fail := len(res) != len(expected)
	if !fail {
		resMap := map[string]bool{}
		for _, w := range res {
			resMap[w] = true
		}
		for _, w := range expected {
			if _, exists := resMap[w]; !exists {
				fail = true
				break
			}
		}
	}
	if fail {
		for _, row := range board {
			t.Logf("[%s]", string(row))
		}
		t.Log("   Words:", words)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestFindWords(t *testing.T) {
	test(
		t,
		[][]byte{
			{'o', 'a', 'a', 'n'},
			{'e', 't', 'a', 'e'},
			{'i', 'h', 'k', 'r'},
			{'i', 'f', 'l', 'v'},
		},
		[]string{"oath", "pea", "eat", "rain"},
		[]string{"eat", "oath"},
	)
	test(
		t,
		[][]byte{
			{'a', 'b'},
			{'c', 'd'},
		},
		[]string{"abcb"},
		[]string{},
	)
}
