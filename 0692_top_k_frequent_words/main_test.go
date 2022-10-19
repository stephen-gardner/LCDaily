package main

import "testing"

func test(t *testing.T, k int, words, expected []string) {
	res := topKFrequent(words, k)
	fail := len(res) != len(expected)
	if !fail {
		for i := range expected {
			if res[i] != expected[i] {
				fail = true
				break
			}
		}
	}
	if fail {
		t.Log("k =", k)
		t.Log("   Words:", words)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestTopKFrequent(t *testing.T) {
	test(
		t, 2,
		[]string{"i", "love", "leetcode", "i", "love", "coding"},
		[]string{"i", "love"},
	)
	test(
		t, 4,
		[]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},
		[]string{"the", "is", "sunny", "day"},
	)
}
