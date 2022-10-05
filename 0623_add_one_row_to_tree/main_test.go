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
		return 1
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

func test(t *testing.T, nums []int, val, depth int, expected []int) {
	tree := treeCreate(nums, 0)
	res := addOneRow(tree, val, depth)
	resArr := make([]*int, treeLenWithNil(tree))
	treeArray(res, resArr, 0)
	fail := false
	for i := range expected {
		if expected[i] == null || resArr[i] == nil {
			if !(expected[i] == null && resArr[i] == nil) {
				fail = true
				break
			}
		} else if *resArr[i] != expected[i] {
			fail = true
			break
		}
	}
	if fail {
		t.Log("    Input:", nums)
		t.Log("      Val:", val)
		t.Log("    Depth:", depth)
		printIntPointerArray(t, resArr, "   Result: ")
		var sb strings.Builder
		sb.WriteByte('[')
		for i := range expected {
			if expected[i] == null {
				sb.WriteRune('•')
			} else {
				sb.WriteString(fmt.Sprintf("%d", expected[i]))
			}
			if i < len(expected)-1 {
				sb.WriteByte(' ')
			}
		}
		sb.WriteByte(']')
		t.Log("Expected:", sb.String())
		t.FailNow()
	}
}

func TestAddOneRow(t *testing.T) {
	test(
		t,
		[]int{4, 2, 6, 3, 1, 5},
		1, 2,
		[]int{4, 1, 1, 2, null, null, 6, 3, 1, null, null, null, null, 5},
	)
	test(
		t,
		[]int{4, 2, null, 3, 1},
		1, 3,
		[]int{4, 2, null, 1, 1, null, null, 3, null, null, 1},
	)
	test(
		t,
		[]int{1, 2, 3, 4},
		5, 4,
		[]int{1, 2, 3, 4, null, null, null, 5, 5},
	)
}
