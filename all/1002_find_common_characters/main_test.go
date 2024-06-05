package main

import (
	"sort"
	"testing"
)

func TestCommonChars(t *testing.T) {
	test := func(words []string, expected []string) {
		res := commonChars(words)
		fail := len(res) != len(expected)
		if !fail {
			sort.Strings(res)
			sort.Strings(expected)
			for i := range expected {
				if expected[i] != res[i] {
					fail = true
					break
				}
			}
		}
		if fail {
			t.Log("Result:", res)
			t.Log("Expected:", expected)
			t.FailNow()
		}
	}
	test([]string{"bella", "label", "roller"}, []string{"e", "l", "l"})
	test([]string{"cool", "lock", "cook"}, []string{"c", "o"})
}
