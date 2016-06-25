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

package set_test

import (
	"fmt"
	"github.com/FabianWe/gocontainer/set"
)

// Iterate over a set of ints and compute the sum of its elements.
func Example() {
	s := set.EmptySet()
	for i := 1; i < 6; i++ {
		s.Add(i)
	}
	sum := 0
	for v := range s.Iter() {
		val := v.(int)
		sum += val
	}
	fmt.Println(sum)
	// Output: 15
}

// Iterate over a set using the Apply function, checks if there is an element
// that is even.
func ExampleApply() {
	s := set.EmptySet()
	for i := 1; i < 20; i += 2 {
		s.Add(i)
	}
	s.Add(2)
	even := false
	f := func(v set.SetValue) bool {
		val := v.(int)
		if val%2 == 0 {
			even = true
			return false
		}
		return true
	}
	s.Apply(f)
	fmt.Println(even)
	// Output: true
}
