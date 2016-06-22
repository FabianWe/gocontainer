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

package queue

import (
	"testing"
)

func checkQueueLen(t *testing.T, q *Queue, len int) {
	n := q.Length()
	if n != len {
		t.Errorf("q.len=%d, wanted %d", n, len)
	}
}

func checkPop(t *testing.T, q *Queue, val int) {
	v := q.Pop()
	if v != val {
		t.Errorf("q.Pop()=%d, wanted %d", v, val)
	}
}

func TestQueue(t *testing.T) {
	q := EmptyQueue()
	checkQueueLen(t, q, 0)
	q.Push(0)
	checkQueueLen(t, q, 1)
	q.Push(1)
	q.Push(2)
	checkQueueLen(t, q, 3)

	checkPop(t, q, 0)
	checkPop(t, q, 1)
	checkQueueLen(t, q, 1)
	checkPop(t, q, 2)
}
