package main

import "testing"

func test(t *testing.T, nums []int, ops []string, expected []interface{}) {
	set := Constructor()
	have := map[int]bool{}
	t.Log("      Ops:", ops)
	t.Log("Arguments:", nums)
	t.Log("Expected:", expected)
	for _, op := range ops {
		switch op {
		case "insert":
			have[nums[0]] = true
			if x := set.Insert(nums[0]); x != expected[0] {
				t.Fatalf("Insert(%d) -> %v != %v (expected)", nums[0], x, expected[0])
			}
			nums = nums[1:]
		case "remove":
			delete(have, nums[0])
			if x := set.Remove(nums[0]); x != expected[0] {
				t.Fatalf("Remove(%d) -> %v != %v (expected)", nums[0], x, expected[0])
			}
			nums = nums[1:]
		case "getRandom":
			if x := set.GetRandom(); !have[x] {
				arr := make([]int, 0, len(have))
				for n := range have {
					arr = append(arr, n)
				}
				t.Log(" Present:", arr)
				t.Fatalf("GetRandom -> %v not in set", x)
				break
			}
		}
		expected = expected[1:]
	}
}

func TestRandomizedSet(t *testing.T) {
	test(t,
		[]int{1, 2, 2, 1, 2},
		[]string{"insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"},
		[]interface{}{true, false, true, 2, true, false, 2})
	test(t,
		[]int{0, 1, 0, 2, 1},
		[]string{"insert", "insert", "remove", "insert", "remove", "getRandom"},
		[]interface{}{true, true, true, true, true, 2})
}
