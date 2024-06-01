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

func test(t *testing.T, nums []int, targetSum int, expected [][]int) {
	tree := treeCreate(nums, 0)
	res := pathSum(tree, targetSum)
	fail := len(res) != len(expected)
	if !fail {
		for i := range expected {
			fail = len(res[i]) != len(expected[i])
			if fail {
				break
			}
			for j := range expected[i] {
				if res[i][j] != expected[i][j] {
					fail = true
					i = len(expected)
					break
				}
			}
		}
	}
	if fail {
		arr := make([]*int, treeLenWithNil(tree))
		treeArray(tree, arr, 0)
		printIntPointerArray(t, arr, "     Input: ")
		t.Log("Target Sum:", targetSum)
		t.Log("    Result:", res)
		t.Log("  Expected:", expected)
		t.FailNow()
	}
}

func TestPathSum(t *testing.T) {
	test(
		t,
		[]int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, null, null, 5, 1},
		22,
		[][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
	)
	test(
		t,
		[]int{1, 2, 3},
		5,
		[][]int{},
	)
	test(
		t,
		[]int{1, 2},
		0,
		[][]int{},
	)
	test(
		t,
		[]int{1},
		1,
		[][]int{{1}},
	)
	test(
		t,
		[]int{},
		0,
		[][]int{},
	)
	test(
		t,
		[]int{1, 2},
		1,
		[][]int{},
	)
}
