gocontainers
============

The Go programming language has no support for generic types and doesn't provide
many datatypes / containers you use very often when programming.
Most of them can be modelled using the existing types, like using a map
to represent a set. However I've found this rather annoying and cumbersome; so I
crated this package.
It currently provides sets (based on a map), stacks (linked list)
and queue (linked list).
This may not be the best way to write your go code but for me that's fine ;).

A set can store anything that can be used as a key for a map, stacks and lists
should work nearly everything.

You can find the documentation on
[godoc.org](https://godoc.org/github.com/FabianWe/gocontainer).

There are also some examples (not all of them are displayed on godoc though)
in the source code, usually in a file called _example_test.go_.

The code is published under the The MIT License (MIT), see file LICENSE.
