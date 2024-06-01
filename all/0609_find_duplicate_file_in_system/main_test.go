package main

import "testing"

func test(t *testing.T, paths []string, expected [][]string) {
	res := findDuplicate(paths)
	fail := len(res) != len(expected)
	if !fail {
		resMap := map[string]bool{}
		for _, paths := range res {
			for _, path := range paths {
				resMap[path] = true
			}
		}
		for i := range expected {
			for j := range expected[i] {
				if _, exists := resMap[expected[i][j]]; !exists {
					fail = true
					i = len(expected)
					break
				}
			}
		}
	}
	if fail {
		t.Log("Inout:", paths)
		t.Log("Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestFindDuplicate(t *testing.T) {
	test(
		t,
		[]string{
			"root/a 1.txt(abcd) 2.txt(efgh)",
			"root/c 3.txt(abcd)",
			"root/c/d 4.txt(efgh)",
			"root 4.txt(efgh)",
		},
		[][]string{
			{"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"},
			{"root/a/1.txt", "root/c/3.txt"},
		},
	)
	test(
		t,
		[]string{
			"root/a 1.txt(abcd) 2.txt(efgh)",
			"root/c 3.txt(abcd)",
			"root/c/d 4.txt(efgh)",
		},
		[][]string{
			{"root/a/2.txt", "root/c/d/4.txt"},
			{"root/a/1.txt", "root/c/3.txt"},
		},
	)
}
