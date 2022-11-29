// Problem: https://leetcode.com/problems/insert-delete-getrandom-o1/
// Results: https://leetcode.com/submissions/detail/851561082/
package main

import (
	"math/rand"
)

type RandomizedSet struct {
	arr  []int
	have map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{have: map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.have[val]; exists {
		return false
	}
	this.arr = append(this.arr, val)
	this.have[val] = len(this.arr) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	i, exists := this.have[val]
	if !exists {
		return false
	}
	swap := this.arr[len(this.arr)-1]
	this.arr[i] = swap
	this.arr = this.arr[:len(this.arr)-1]
	this.have[swap] = i
	delete(this.have, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	i := rand.Intn(len(this.arr))
	return this.arr[i]
}
