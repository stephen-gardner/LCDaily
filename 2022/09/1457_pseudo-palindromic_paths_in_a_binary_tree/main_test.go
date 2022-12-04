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

func test(t *testing.T, nums []int, expected int) {
	tree := treeCreate(nums, 0)
	res := pseudoPalindromicPaths(tree)
	if res != expected {
		arr := make([]*int, treeLenWithNil(tree))
		treeArray(tree, arr, 0)
		printIntPointerArray(t, arr, "Input: ")
		t.Fatalf("%d != %d (expected)", res, expected)
	}
}

func TestPseudoPalindromicPaths(t *testing.T) {
	test(t, []int{2, 3, 1, 3, 1, null, 1}, 2)
	test(t, []int{2, 1, 1, 1, 3, null, null, null, null, null, 1}, 1)
	test(t, []int{9}, 1)
	test(t, []int{5, null, 7}, 0)
}
