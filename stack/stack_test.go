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

package stack

import (
	"testing"
)

func checkStackLen(t *testing.T, s *Stack, len int) {
	n := s.Length()
	if n != len {
		t.Errorf("s.len=%d, wanted %d", n, len)
	}
}

func checkPop(t *testing.T, s *Stack, val int) {
	v := s.Pop()
	if v != val {
		t.Errorf("s.Pop()=%d, wanted %d", v, val)
	}
}

func TestStack(t *testing.T) {
	s := EmptyStack()
	checkStackLen(t, s, 0)
	s.Push(0)
	checkStackLen(t, s, 1)
	s.Push(1)
	s.Push(2)
	checkStackLen(t, s, 3)

	checkPop(t, s, 2)
	checkPop(t, s, 1)
	checkStackLen(t, s, 1)
	checkPop(t, s, 0)
}
