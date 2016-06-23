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

type SetValue interface{}

const maxBuffer = 10000

// Set is a unordered collection of unique values.
type Set struct {
	elements map[SetValue]struct{}
}

func EmptySet() *Set {
	return &Set{make(map[SetValue]struct{})}
}

func (this *Set) Length() int {
	return len(this.elements)
}

func (this *Set) Empty() bool {
	return this.Length() == 0
}

func (this *Set) Add(val SetValue) {
	this.elements[val] = struct{}{}
}

func (this *Set) Remove(val SetValue) {
	delete(this.elements, val)
}

func (this *Set) Contains(val SetValue) bool {
	_, res := this.elements[val]
	return res
}

func (this *Set) Subset(other *Set) bool {
	for val := range this.Iter() {
		if !other.Contains(val) {
			return false
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

// Return a chanel to iterate over all elements in the set.
// Important: Only use this iterator if you plan to iterate over *all* objects,
// so don't break the for loop with a break or something like that.
// If you plan to break the loop you can use Apply.
func (this *Set) Iter() <-chan SetValue {
	bufferSize := this.Length()
	if bufferSize > maxBuffer {
		bufferSize = maxBuffer
	}
	ch := make(chan SetValue, bufferSize)
	go func() {
		for k, _ := range this.elements {
			ch <- k
		}
		close(ch)
	}()
	return ch
}

// Applies a function to all values (acts like an iterator).
// The function argument is applied to all values in the set, if the function
// returns true the next element will be visited, otherwise the iteration
// will stop.
// See the example.
func (this *Set) Apply(f func(val SetValue) bool) {
	for k, _ := range this.elements {
		if !f(k) {
			break
		}
	}
}

func (this *Set) Pop() (SetValue, bool) {
	var val SetValue = nil
	succ := false
	for k, _ := range this.elements {
		val = k
		succ = true
		break
	}
	if succ {
		this.Remove(val)
	}
	return val, succ
}

func (this *Set) Copy() *Set {
	res := EmptySet()
	for k, _ := range this.elements {
		res.Add(k)
	}
	return res
}
