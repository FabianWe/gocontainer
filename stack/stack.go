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

// Package provides a simple LIFO (last-in-first-out) datatype.
package stack

// Type for elements stored in the stack nodes.
type StackValue interface{}

// A node has a value it stores and a pointer to the next node (nil if there is
// no next node).
type stackNode struct {
	value StackValue
	next  *stackNode
}

// Stack is a simple LIFO (last-in-first-out) datatype.
type Stack struct {
	first *stackNode
	size  int
}

// EmptyStack creates a new stack without any elements.
func EmptyStack() *Stack {
	return &Stack{nil, 0}
}

// Length returns the number of elements currently on the stack.
func (this *Stack) Length() int {
	return this.size
}

// Empty returns true if there are no elements on the stack.
func (this *Stack) Empty() bool {
	return this.Length() == 0
}

// Push adds a new value to the front of the stack.
func (this *Stack) Push(val StackValue) {
	this.first = &stackNode{val, this.first}
	this.size++
}

// Pop returns the first element of the stack, nil if the stack is empty.
func (this *Stack) Pop() StackValue {
	if this.Empty() {
		return nil
	}
	v := this.first.value
	this.first = this.first.next
	this.size--
	return v
}
