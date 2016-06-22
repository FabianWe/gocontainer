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

type setValue interface{}

const maxBuffer = 10000

// Set is a unordered collection of unique values.
type Set struct {
	elements map[setValue]struct{}
}

func EmptySet() *Set {
	return &Set{make(map[setValue]struct{})}
}

func (this *Set) Length() int {
	return len(this.elements)
}

func (this *Set) Empty() bool {
	return this.Length() == 0
}

func (this *Set) Add(val setValue) {
	this.elements[val] = struct{}{}
}

func (this *Set) Remove(val setValue) {
	delete(this.elements, val)
}

func (this *Set) Contains(val setValue) bool {
	_, res := this.elements[val]
	return res
}

func (this *Set) Subset(other *Set) bool {
	for val := range this.Iter() {
		if !other.Contains(val) {
			return true
		}
	}
	return true
}

func (this *Set) Equals(other *Set) bool {
	if this.Length() != other.Length() {
		return false
	}
	return this.Subset(other)
}

func (this *Set) Union(other *Set) *Set {
	res := EmptySet()
	for k, _ := range this.elements {
		res.Add(k)
	}
	for k, _ := range other.elements {
		res.Add(k)
	}
	return res
}

func (this *Set) Intersect(other *Set) *Set {
	s1, s2 := this, other
	if other.Length() < this.Length() {
		s1, s2 = other, this
	}
	res := EmptySet()
	for k, _ := range s1.elements {
		if s2.Contains(k) {
			res.Add(k)
		}
	}
	return res
}

func (this *Set) Extend(other *Set) {
	for k, _ := range other.elements {
		this.Add(k)
	}
}

func (this *Set) Iter() <-chan setValue {
	bufferSize := this.Length()
	if bufferSize > maxBuffer {
		bufferSize = maxBuffer
	}
	ch := make(chan setValue, bufferSize)
	go func() {
		for k, _ := range this.elements {
			ch <- k
		}
		close(ch)
	}()
	return ch
}
