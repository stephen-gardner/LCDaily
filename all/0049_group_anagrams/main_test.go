package main

import "testing"

func test(t *testing.T, strs []string, expected [][]string) {
	resArr := groupAnagrams(strs)
	res := map[string]map[string]bool{}
	for _, arr := range resArr {
		set := map[string]bool{}
		for _, s := range arr {
			set[s] = true
			res[s] = set
		}
	}
	for _, set := range expected {
		for _, s := range set {
			if !res[s][s] {
				t.Log("   Input:", strs)
				t.Log("  Result:", resArr)
				t.Log("Expected:", expected)
				t.FailNow()
			}
		}
	}
}

func TestGroupAnagrams(t *testing.T) {
	test(
		t,
		[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
		[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
	)
	test(
		t,
		[]string{""},
		[][]string{{""}},
	)
	test(
		t,
		[]string{"a"},
		[][]string{{"a"}},
	)
	test(
		t,
		[]string{"eat", "tea", "tan", "ate", "nat", "bat", "ac", "bd", "aac", "bbd", "aacc", "bbdd", "acc", "bdd"},
		[][]string{{"tan", "nat"}, {"bbdd"}, {"acc"}, {"aac"}, {"bbd"}, {"aacc"}, {"bdd"}, {"eat", "tea", "ate"}, {"bat"}, {"ac"}, {"bd"}},
	)
}
