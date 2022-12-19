package main

import "testing"

func test(t *testing.T, nums []int, ops []string, expected []interface{}) {
	var queue MyQueue
	res := []interface{}{}
	i := 0
	for _, op := range ops {
		switch op {
		case "MyQueue":
			queue = Constructor()
			res = append(res, nil)
		case "push":
			queue.Push(nums[i])
			res = append(res, nil)
			i++
		case "pop":
			res = append(res, queue.Pop())
		case "peek":
			res = append(res, queue.Peek())
		case "empty":
			res = append(res, queue.Empty())
		}
	}
	for i := range expected {
		if res[i] != expected[i] {
			t.Log("    Nums:", nums)
			t.Log("     Ops:", ops)
			t.Log("  Result:", res)
			t.Log("Expected:", expected)
			t.FailNow()
		}
	}
}

func TestMyQueue(t *testing.T) {
	test(t,
		[]int{1, 2},
		[]string{"MyQueue", "push", "push", "peek", "pop", "empty"},
		[]interface{}{nil, nil, nil, 1, 1, false})
}
