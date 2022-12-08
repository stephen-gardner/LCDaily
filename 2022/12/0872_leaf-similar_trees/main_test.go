package main

import (
	"math"
	"testing"

	"github.com/stephen-gardner/LCDaily/test/tree"
)

const null = math.MaxInt

func TestGetSequence(t *testing.T) {
	seq := []byte{}
	//        3
	//    5       1
	//  6   2   9   8
	// N N 7 4 N N N N
	expected := []byte{6, 7, 4, 9, 8}
	testTree := tree.Create([]int{3, 5, 1, 6, 2, 9, 8, null, null, 7, 4}, 0)
	getSequence(&seq, testTree)
	fail := len(seq) != len(expected)
	if !fail {
		for i := range expected {
			if seq[i] != expected[i] {
				fail = true
				break
			}
		}
	}
	if fail {
		t.Log("   Input:", testTree.String())
		t.Log("  Result:", seq)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func test(t *testing.T, nums1, nums2 []int, expected bool) {
	tree1 := tree.Create(nums1, 0)
	tree2 := tree.Create(nums2, 0)
	if leafSimilar(tree1, tree2) != expected {
		t.Log("Tree 1:", tree1.String())
		t.Log("Tree 2:", tree2.String())
		t.Fatalf("%v != %v (expected)", !expected, expected)
	}
}

func TestLeafSimilar(t *testing.T) {
	test(t,
		[]int{3, 5, 1, 6, 2, 9, 8, null, null, 7, 4},
		[]int{3, 5, 1, 6, 7, 4, 2, null, null, null, null, null, null, 9, 8},
		true)
	test(t,
		[]int{1, 2, 3},
		[]int{1, 3, 2},
		false)
	test(t,
		[]int{3, 5, 1, 6, 7, 4, 2, null, null, null, null, null, null, 9, 11, null, null, 8, 10},
		[]int{3, 5, 1, 6, 2, 9, 8, null, null, 7, 4},
		false)
}
