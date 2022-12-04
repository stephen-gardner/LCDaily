package main

import "testing"

func test(t *testing.T, ops [][]interface{}, expected []string) {
	tm := Constructor()
	res := []string{}
	for _, op := range ops {
		switch len(op) {
		case 3:
			tm.Set(op[0].(string), op[1].(string), op[2].(int))
		case 2:
			res = append(res, tm.Get(op[0].(string), op[1].(int)))
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

func TestTimeMap(t *testing.T) {
	test(
		t,
		[][]interface{}{
			{"foo", "bar", 1},
			{"foo", 1},
			{"foo", 3},
			{"foo", "bar2", 4},
			{"foo", 4},
			{"foo", 5},
		},
		[]string{
			"bar",
			"bar",
			"bar2",
			"bar2",
		},
	)
}
