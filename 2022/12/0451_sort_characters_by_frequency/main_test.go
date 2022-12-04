package main

import "testing"

func test(t *testing.T, s, expected string) {
	if res := frequencySort(s); res != expected {
		t.Log("   Input:", s)
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestFrequencySort(t *testing.T) {
	test(t, "tree", "eetr")
	test(t, "cccaaa", "aaaccc")
	test(t, "Aabb", "bbAa")
	test(t, "2a554442f544asfasssffffasss", "sssssssffffff44444aaaa55522")
}
