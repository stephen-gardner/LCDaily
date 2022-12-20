// Problem: https://leetcode.com/problems/keys-and-rooms/
// Results: https://leetcode.com/problems/keys-and-rooms/submissions/862621010/
package main

func canVisitAllRooms(rooms [][]int) bool {
	queue := []int{0}
	unlocked := map[int]bool{0: true}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			for _, k := range rooms[queue[i]] {
				if !unlocked[k] {
					unlocked[k] = true
					queue = append(queue, k)
				}
			}
		}
		queue = queue[size:]
	}
	for i := range rooms {
		if !unlocked[i] {
			return false
		}
	}
	return true
}
