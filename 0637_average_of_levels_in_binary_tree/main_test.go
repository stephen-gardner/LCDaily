package main

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

const null = math.MaxInt

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func treeCreate(nums []int, pos int) *TreeNode {
	if pos >= len(nums) || nums[pos] == null {
		return nil
	}
	return &TreeNode{
		Val:   nums[pos],
		Left:  treeCreate(nums, (2*pos)+1),
		Right: treeCreate(nums, (2*pos)+2),
	}
}

func treeArray(root *TreeNode, nums []*int, pos int) {
	if pos >= len(nums) || root == nil {
		return
	}
	nums[pos] = &root.Val
	treeArray(root.Left, nums, (pos*2)+1)
	treeArray(root.Right, nums, (pos*2)+2)
}

func treeHeight(curr *TreeNode) int {
	if curr == nil {
		return 0
	}
	return 1 + max(treeHeight(curr.Left), treeHeight(curr.Right))
}

func treeLenWithNil(root *TreeNode) int {
	length := 1
	height := treeHeight(root)
	for i := 0; i < height; i++ {
		length += (2 * i)
	}
	return length
}

func printIntPointerArray(t *testing.T, arr []*int, prefix string) {
	sb := strings.Builder{}
	sb.WriteRune('[')
	for i := 0; i < len(arr); i++ {
		if arr[i] == nil {
			sb.WriteRune('•')
		} else {
			sb.WriteString(fmt.Sprintf("%d", *arr[i]))
		}
		if i < len(arr)-1 {
			sb.WriteRune(' ')
		}
	}
	sb.WriteRune(']')
	t.Logf("%s%s", prefix, sb.String())
}

func test(t *testing.T, nums []int, expected []float64) {
	tree := treeCreate(nums, 0)
	res := averageOfLevels(tree)
	fail := len(res) != len(expected)
	if !fail {
		for i := 0; i < len(res); i++ {
			if res[i] != expected[i] {
				fail = true
				break
			}
		}
	}
	if fail {
		arr := make([]*int, treeLenWithNil(tree))
		treeArray(tree, arr, 0)
		printIntPointerArray(t, arr, "   Input: ")
		t.Log("  Result:", res)
		t.Log("Expected:", expected)
		t.FailNow()
	}
}

func TestAverageOfLevels(t *testing.T) {
	test(
		t,
		[]int{3, 9, 20, null, null, 15, 7},
		[]float64{3.00000, 14.50000, 11.00000},
	)

	test(
		t,
		[]int{3, 9, 20, 15, 7},
		[]float64{3.00000, 14.50000, 11.00000},
	)
}
