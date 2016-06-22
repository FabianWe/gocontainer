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

package queue_test

import (
	"fmt"
	"github.com/FabianWe/gocontainer/queue"
)

func Example() {
	q := queue.EmptyQueue()
	q.Push(0)
	q.Push(1)
	q.Push(2)

	first := q.Pop()
	second := q.Pop()
	third := q.Pop()

	fmt.Println(first, second, third)
	// Output: 0 1 2
}

func ExampleLoop() {
	q := queue.EmptyQueue()
	for i := 1; i < 6; i++ {
		q.Push(i)
	}
	sum := 0
	for !q.Empty() {
		var val int = q.Pop().(int)
		sum += val
	}
	fmt.Println(sum)
	// Output: 15
}
