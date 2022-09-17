package main

import "testing"

func test(t *testing.T, words []string, expected [][]int) {
	res := palindromePairs(words)
	fail := len(res) != len(expected)
	if !fail {
		for i := range expected {
			if res[i][0] != expected[i][0] || res[i][1] != expected[i][1] {
				fail = true
				break
			}
		}
	}
	if fail {
		t.Log("   Input:", words)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestPalindromePairs(t *testing.T) {
	test(t, []string{"abcd", "dcba", "lls", "s", "sssll"}, [][]int{{0, 1}, {1, 0}, {2, 4}, {3, 2}})
	test(t, []string{"bat", "tab", "cat"}, [][]int{{0, 1}, {1, 0}})
	test(t, []string{"a", ""}, [][]int{{0, 1}, {1, 0}})
	test(t, []string{"a", "abc", "aba", ""}, [][]int{{0, 3}, {2, 3}, {3, 0}, {3, 2}})
	test(t, []string{"a", "b", "c", "ab", "ac", "aa"}, [][]int{{0, 5}, {1, 3}, {2, 4}, {3, 0}, {4, 0}, {5, 0}})
	test(t, []string{"tabbbbb", "bat"}, [][]int{{0, 1}})
	test(t, []string{"aaaa"}, [][]int{})
}
