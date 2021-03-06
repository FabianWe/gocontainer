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

func checkSubset(t *testing.T, s1, s2 *Set, subset bool) {
	res := s1.Subset(s2)
	if res != subset {
		if subset {
			t.Error("Expected s1 to be a subset of s2, but it is not")
		} else {
			t.Error("Expected s1 no to be a subset of s2, but it is")
		}
	}
}

func checkSetEqual(t *testing.T, s1, s2 *Set, equal bool) {
	res := s1.Equals(s2)
	if res != equal {
		if equal {
			t.Error("Expected sets to be equal, but they're not")
		} else {
			t.Error("Expected sets to be unequal, but they're equal")
		}
	}
}

func TestSetCompare(t *testing.T) {
	s1 := EmptySet()
	s2 := EmptySet()
	checkSetEqual(t, s1, s2, true)
	checkSetEqual(t, s2, s1, true)
	checkSubset(t, s1, s2, true)
	s2.Add(0)
	checkSubset(t, s1, s2, true)
	checkSubset(t, s2, s1, false)
	for i := 1; i < 5; i++ {
		s1.Add(i)
		s2.Add(i)
	}
	checkSubset(t, s1, s2, true)
	checkSubset(t, s2, s1, false)
}

func TestSetUnion(t *testing.T) {
	s1 := EmptySet()
	s2 := EmptySet()

	for i := 0; i < 5; i++ {
		s1.Add(i)
	}
	for i := 2; i < 10; i++ {
		s2.Add(i)
	}

	resSet := EmptySet()
	for i := 0; i < 10; i++ {
		resSet.Add(i)
	}

	if !s1.Union(s2).Equals(resSet) {
		t.Error("Union didn't work")
	}
}

func TestSetIntersect(t *testing.T) {
	s1 := EmptySet()
	s2 := EmptySet()

	for i := 0; i < 5; i++ {
		s1.Add(i)
	}
	for i := 2; i < 10; i++ {
		s2.Add(i)
	}

	resSet := EmptySet()
	for i := 2; i < 5; i++ {
		resSet.Add(i)
	}

	if !s1.Intersect(s2).Equals(resSet) {
		t.Error("Intersect didn't work")
	}
}
