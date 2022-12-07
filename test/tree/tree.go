package tree

import (
	"math"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const null = math.MaxInt

func Create(nums []int, pos int) *TreeNode {
	if pos >= len(nums) || nums[pos] == null {
		return nil
	}
	return &TreeNode{
		Val:   nums[pos],
		Left:  Create(nums, (2*pos)+1),
		Right: Create(nums, (2*pos)+2),
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (tree *TreeNode) Height() int {
	if tree == nil {
		return 0
	}
	return 1 + max(tree.Left.Height(), tree.Right.Height())
}

func toArray(root *TreeNode, nums []*int, pos int) {
	if pos >= len(nums) || root == nil {
		return
	}
	nums[pos] = &root.Val
	toArray(root.Left, nums, (pos*2)+1)
	toArray(root.Right, nums, (pos*2)+2)
}

func (tree *TreeNode) String() string {
	size := (1 << tree.Height()) - 1
	arr := make([]*int, size)
	toArray(tree, arr, 0)
	var sb strings.Builder
	sb.WriteByte('[')
	for i := 0; i < len(arr); i++ {
		if arr[i] == nil {
			sb.WriteRune('•')
		} else {
			sb.WriteString(strconv.Itoa(*arr[i]))
		}
		if i < len(arr)-1 {
			sb.WriteByte(' ')
		}
	}
	sb.WriteByte(']')
	return sb.String()
}
