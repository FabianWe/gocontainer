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

package set

import (
  "testing"
)

func checkSetLen(t *testing.T, s *Set, len int) {
  n := s.Length()
  if n != len {
    t.Errorf("s.len=%d, wanted %d", n, len)
  }
}

func checkContains(t *testing.T, s *Set, val int, res bool) {
  if s.Contains(val) != res {
    if res {
      t.Errorf("Expected that %d is contained in set, but it isn't", val)
    } else {
      t.Errorf("Expected that %d isn't contained in set, but it is", val)
    }
  }
}

func TestSet(t *testing.T) {
  s := EmptySet()
  checkSetLen(t, s, 0)
  checkContains(t, s, 0, false)
  s.Add(0)
  s.Add(0)
  checkContains(t, s, 0, true)
  checkSetLen(t, s, 1)
  for i := 1; i < 10; i++ {
    s.Add(i)
  }
  checkSetLen(t, s, 10)
  checkContains(t, s, 9, true)
  checkContains(t, s, 11, false)
  s.Remove(9)
  checkContains(t, s, 9, false)
}
