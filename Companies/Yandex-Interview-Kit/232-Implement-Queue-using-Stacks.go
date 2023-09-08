package main

type MyQueue struct {
	size int
	head *Node // MRU what to return or peek
	tail *Node // LRU where to append
}

type Node struct {
	value int
	left  *Node
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	elem := &Node{value: x}

	if this.size == 0 {
		this.head, this.tail = elem, elem
	} else {
		this.tail.left = elem
		this.tail = elem
	}

	this.size++
}

func (this *MyQueue) Pop() int {
	if !this.Empty() {
		elem := this.head
		this.head = this.head.left
		if this.head == nil {
			this.tail = nil
		}

		this.size--
		return elem.value
	}
	return 0
}

func (this *MyQueue) Peek() int {
	return this.head.value
}

func (this *MyQueue) Empty() bool {
	return this.size == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
