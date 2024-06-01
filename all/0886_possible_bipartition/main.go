// Problem: https://leetcode.com/problems/possible-bipartition/
// Results: https://leetcode.com/problems/possible-bipartition/submissions/863173889/
package main

const (
	Unassigned = -1
	Red        = 0
	Blue       = 1
)

type Member struct {
	opposed []*Member
	team    int8
}

func hasConflict(curr *Member, expected int8) bool {
	if curr.team == Unassigned {
		curr.team = expected
	} else {
		return curr.team != expected
	}
	for _, next := range curr.opposed {
		if hasConflict(next, 1-expected) {
			return true
		}
	}
	return false
}

func possibleBipartition(n int, dislikes [][]int) bool {
	people := make([]*Member, n+1)
	for i := 1; i <= n; i++ {
		people[i] = &Member{team: Unassigned}
	}
	for _, data := range dislikes {
		people[data[0]].opposed = append(people[data[0]].opposed, people[data[1]])
		people[data[1]].opposed = append(people[data[1]].opposed, people[data[0]])
	}
	for i := 1; i <= n; i++ {
		origin := people[i]
		if origin.team == Unassigned && hasConflict(origin, Red) {
			return false
		}
	}
	return true
}
