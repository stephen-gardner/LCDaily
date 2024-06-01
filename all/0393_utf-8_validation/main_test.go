package main

import "testing"

func test(t *testing.T, data []int, expected bool) {
	res := validUtf8(data)
	if res != expected {
		t.Log("Input:", data)
		t.Fatalf("%v != %v (expected)", res, expected)
	}
}

func TestValidUtf8(t *testing.T) {
	test(t, []int{197, 130, 1}, true)
	test(t, []int{235, 140, 4}, false)
	test(t, []int{240, 162, 138, 147}, true)
}
