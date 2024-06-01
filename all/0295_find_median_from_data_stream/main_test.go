package main

import "testing"

func test(t *testing.T, ops []string, nums []int, expected []float64) {
	res := []float64{}
	f := Constructor()
	for _, op := range ops {
		switch op {
		case "addNum":
			f.AddNum(nums[0])
			nums = nums[1:]
		case "findMedian":
			res = append(res, f.FindMedian())
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
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestMedianFinder(t *testing.T) {
	test(
		t,
		[]string{"addNum", "addNum", "findMedian", "addNum", "findMedian"},
		[]int{1, 2, 3},
		[]float64{1.5, 2.0},
	)
	test(
		t,
		[]string{
			"addNum", "findMedian", "addNum", "findMedian", "addNum",
			"findMedian", "addNum", "findMedian", "addNum", "findMedian",
			"addNum", "findMedian", "addNum", "findMedian", "addNum",
			"findMedian", "addNum", "findMedian", "addNum", "findMedian",
			"addNum", "findMedian",
		},
		[]int{6, 10, 2, 6, 5, 0, 6, 3, 1, 0, 0},
		[]float64{6.0, 8.0, 6.0, 6.0, 6.0, 5.5, 6.0, 5.5, 5.0, 4.0, 3.0},
	)
}
