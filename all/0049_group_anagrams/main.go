package main

func groupAnagrams(strs []string) [][]string {
	counts := map[int][]string{}
	counter := [26]byte{}
	for _, s := range strs {
		for _, c := range s {
			counter[c-'a']++
		}
		// djb2: http://www.cse.yorku.ca/~oz/hash.html
		hash := 5381
		for i := range counter {
			hash = (hash * 33) ^ int(counter[i])
			counter[i] = 0
		}
		counts[hash] = append(counts[hash], s)
	}
	res := [][]string{}
	for _, set := range counts {
		res = append(res, set)
	}
	return res
}
