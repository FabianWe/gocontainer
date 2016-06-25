// The MIT License (MIT)

// Copyright (c) 2016 Fabian Wenzelmann

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package provides a simple FIFO (first-in-first-out) datatype.
package queue

// Type for elements stored in the queue nodes.
type QueueValue interface{}

// A node has a value it stores and a pointer to the next node (nil if there is
// no next node).
type queueNode struct {
	value QueueValue
	next  *queueNode
}

// Stack is a simple FIFO (first-in-first-out) datatype.
type Queue struct {
	first *queueNode
	last  *queueNode
	size  int
}

// EmptyQueue creates a new queue without any elements.
func EmptyQueue() *Queue {
	return &Queue{nil, nil, 0}
}

// Length returns the number of elements currently in the queue.
func (this *Queue) Length() int {
	return this.size
}

// Empty returns true if there are no elements in the queue.
func (this *Queue) Empty() bool {
	return this.Length() == 0
}

// Push adds a new value to the back of the queue.
func (this *Queue) Push(val QueueValue) {
	n := queueNode{val, nil}
	if this.Empty() {
		// first and last are the same new node
		this.first = &n
		this.last = &n
	} else {
		// simply append to the last node
		// and update last node
		this.last.next = &n
		this.last = &n
	}
	this.size++
}

// Pop returns the first element of the queue, nil if the queue is empty.
func (this *Queue) Pop() QueueValue {
	if this.Empty() {
		return nil
	}
	next := this.first
	if this.Length() == 1 {
		this.first = nil
		this.last = nil
	} else {
		this.first = this.first.next
	}
	this.size--
	return next.value
}
