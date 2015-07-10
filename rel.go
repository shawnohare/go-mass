// Package rel provides interfaces and functions that act on lists
// of objects that have some embedding into the real numbers.
// A common use case is to employ a single data structure modeling
// a collection of (object, value) pairs that handles sorting, partitioning,
// and statistical calculations over the corresponding set of real values.
package rel

// A List is a type that can sorted/partitioned/fetched from and
// represents an ordered set of objects that can be mapped to a real number.
type List interface {
	// Len is the number of elements in a the list
	Len() int
	// Get the ith element.
	Get(i int) interface{}
	// Get the ith element's associated real value.
	Val(i int) float64
}
