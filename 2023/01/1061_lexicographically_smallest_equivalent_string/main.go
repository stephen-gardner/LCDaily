// Problem: https://leetcode.com/problems/lexicographically-smallest-equivalent-string/
// Results: https://leetcode.com/problems/lexicographically-smallest-equivalent-string/submissions/878020051/
package main

func findSmallest(equiv [][]bool, visited []bool, key int) int {
	visited[key] = true
	smallest := key
	for i, isEquivalent := range equiv[key] {
		if !isEquivalent || visited[i] {
			continue
		}
		if res := findSmallest(equiv, visited, i); res < smallest {
			smallest = res
		}
	}
	return smallest
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	equiv := make([][]bool, 26)
	for i := range equiv {
		equiv[i] = make([]bool, 26)
	}
	for i := range s1 {
		c1, c2 := s1[i]-'a', s2[i]-'a'
		equiv[c2][c1] = true
		equiv[c1][c2] = true
	}
	smallest := [26]byte{}
	for i := 0; i < 26; i++ {
		smallest[i] = byte('a' + findSmallest(equiv, make([]bool, 26), i))
	}
	res := []byte(baseStr)
	for i := range res {
		res[i] = smallest[res[i]-'a']
	}
	return string(res)
}
