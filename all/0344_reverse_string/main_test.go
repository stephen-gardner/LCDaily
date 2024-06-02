package main

import "testing"

func test(t *testing.T, input, expected string) {
	res := []byte(input)
	reverseString(res)
	if res := string(res); res != expected {
		t.Fatalf("\nInput: %s\nOutput: %s\nExpected: %s", input, res, expected)
	}
}

func TestReverseString(t *testing.T) {
	test(t, "hello", "olleh")
	test(t, "Hannah", "hannaH")
}
