package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func Array(curr *ListNode) []int {
	arr := []int{}
	for curr != nil {
		arr = append(arr, curr.Val)
		curr = curr.Next
	}
	return arr
}

func Create(nums []int) *ListNode {
	var head *ListNode
	curr := &head
	for _, n := range nums {
		*curr = &ListNode{Val: n}
		curr = &(*curr).Next
	}
	return head
}
