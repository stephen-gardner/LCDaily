package main

import "testing"

func test(t *testing.T, k int, nums []int, ops []string, expected []interface{}) {
	q := Constructor(k)
	res := []interface{}{}
	for _, op := range ops {
		switch op {
		case "enQueue":
			res = append(res, q.EnQueue(nums[0]))
			nums = nums[1:]
		case "deQueue":
			res = append(res, q.DeQueue())
		case "Front":
			res = append(res, q.Front())
		case "Rear":
			res = append(res, q.Rear())
		case "isEmpty":
			res = append(res, q.IsEmpty())
		case "isFull":
			res = append(res, q.IsFull())
		default:
			break
		}
	}
	fail := len(res) != len(expected)
	if !fail {
		for i := range expected {
			if res[i] != expected[i] {
				fail = true
				break
			}
		}
	}
	if fail {
		t.Log("Operations:", ops)
		t.Log("    Result:", res)
		t.Log("  Expected:", expected)
		t.FailNow()
	}
}

func TestMyCircularQueue(t *testing.T) {
	test(t, 3,
		[]int{1, 2, 3, 4, 4},
		[]string{
			"enQueue", "enQueue", "enQueue", "enQueue", "Rear", "isFull",
			"deQueue", "enQueue", "Rear", "isEmpty", "Front", "deQueue",
			"deQueue", "deQueue", "deQueue", "isFull", "isEmpty", "Front",
			"Rear",
		},
		[]interface{}{
			true, true, true, false, 3, true, true, true, 4, false, 2, true,
			true, true, false, false, true, -1, -1,
		},
	)
}
