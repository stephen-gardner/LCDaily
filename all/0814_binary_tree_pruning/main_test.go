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

func test(t *testing.T, nums []int, expected []int) {
	tree := treeCreate(nums, 0)
	inputArr := make([]*int, treeLenWithNil(tree))
	treeArray(tree, inputArr, 0)
	res := pruneTree(tree)
	resArr := make([]*int, treeLenWithNil(tree))
	treeArray(res, resArr, 0)
	expectedTree := treeCreate(expected, 0)
	expectedArr := make([]*int, treeLenWithNil(expectedTree))
	treeArray(expectedTree, expectedArr, 0)
	fail := len(resArr) != len(expectedArr)
	if !fail {
		for i := 0; i < len(resArr) && !fail; i++ {
			if resArr[i] == nil || expectedArr[i] == nil {
				fail = resArr[i] != expectedArr[i]
			} else {
				fail = *resArr[i] != *expectedArr[i]
			}
		}
	}
	if fail {
		printIntPointerArray(t, inputArr, "   Input: ")
		printIntPointerArray(t, resArr, "  Result: ")
		printIntPointerArray(t, expectedArr, "Expected: ")
		t.FailNow()
	}
}

func TestPruneTree(t *testing.T) {
	test(
		t,
		[]int{1, null, 0, null, null, 0, 1},
		[]int{1, null, 0, null, null, null, 1},
	)
	test(
		t,
		[]int{1, 0, 1, 0, 0, 0, 1},
		[]int{1, null, 1, null, null, null, 1},
	)
	test(
		t,
		[]int{1, 1, 0, 1, 1, 0, 1, 0},
		[]int{1, 1, 0, 1, 1, null, 1},
	)
	test(
		t,
		[]int{0},
		[]int{},
	)
	test(
		t,
		[]int{1},
		[]int{1},
	)
}
