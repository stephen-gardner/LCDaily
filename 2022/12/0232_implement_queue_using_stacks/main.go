// Problem: https://leetcode.com/problems/implement-queue-using-stacks/
// Results: https://leetcode.com/problems/implement-queue-using-stacks/submissions/861841413/
package main

type (
	Stack   []int
	MyQueue struct {
		read  Stack
		write Stack
	}
)

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	stack := *s
	x := stack[len(stack)-1]
	*s = stack[:len(stack)-1]
	return x
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func Constructor() MyQueue {
	return MyQueue{
		read:  Stack{},
		write: Stack{},
	}
}

func (this *MyQueue) Push(x int) {
	for !this.read.IsEmpty() {
		this.write.Push(this.read.Pop())
	}
	this.write.Push(x)
}

func (this *MyQueue) Pop() int {
	for !this.write.IsEmpty() {
		this.read.Push(this.write.Pop())
	}
	return this.read.Pop()
}

func (this *MyQueue) Peek() int {
	for !this.write.IsEmpty() {
		this.read.Push(this.write.Pop())
	}
	return this.read.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.read.IsEmpty() && this.write.IsEmpty()
}
