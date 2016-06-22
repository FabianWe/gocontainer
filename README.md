gocontainers
============

The Go programming language has no support for generic types and doesn't provide
many datatypes / containers you use very often when programming.
Most of them can be modelled using the existing types, like using a map
to represent a set. However I've found this rather annoying and so crated this
package. It currently provides sets (based on a map), stacks (linked list)
and queue (linked list).

A set can store anything that can be used as a key for a map, stacks and lists
should work nearly everything.

More explanation to come!
