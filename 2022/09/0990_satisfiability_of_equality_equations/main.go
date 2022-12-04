// Problem: https://leetcode.com/problems/satisfiability-of-equality-equations/
// Results: https://leetcode.com/submissions/detail/808739719/
package main

func propagate(equiv *[26][26]bool, n1, n2 int) {
	if equiv[n1][n2] == true {
		return
	}
	equiv[n1][n2] = true
	for n, equal := range equiv[n1] {
		if equal {
			propagate(equiv, n, n2)
			propagate(equiv, n2, n)
		}
	}
}

func equationsPossible(equations []string) bool {
	equiv := [26][26]bool{}
	for _, eq := range equations {
		if eq[1] != '=' {
			if eq[0] == eq[3] {
				return false
			}
			continue
		}
		n1, n2 := int(eq[0]-'a'), int(eq[3]-'a')
		propagate(&equiv, n1, n2)
		propagate(&equiv, n2, n1)
	}
	for _, eq := range equations {
		if eq[1] != '!' {
			continue
		}
		n1, n2 := eq[0]-'a', eq[3]-'a'
		if equiv[n1][n2] || equiv[n2][n1] {
			return false
		}
	}
	return true
}
