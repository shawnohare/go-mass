// Package mass provides interfaces and function to perform
// set and statistical  operations on data that has some notion of mass
// norm.  Examples of such data include:
//
// - planets and their masses.
// - functions and their norms.
// - a collection of random variables and an associated response variable.
//
// A typical structure implementing the mass.Interface is a slice
// []A, where A is a type with some attribute corresponding to a real number.
// This attribute then induces a total ordering on []A that we might wish to:
//
// - Sort by.
// - Partition by.
// - Select subsets by.
package mass

import (
	"sort"
)

// TODO consider removing dependency on sort interface.
// We will never use this explicitly if we demand the data has been
// previously sorted.

// A type, typically a collection, that satisfies mass.Interface
// can be partitioned/sorted/selected from by the functions in this package.
// The  methods require that the elements can be enumerated by an integer index
// and assigned an numeric value.
type Interface interface {
	// Len is the number of elements in a the collection
	Len() int
	// Less reports whether the element with index i should sort before
	// element with index j.  It induces a total ordering on the collection.
	// Typically Less(i, j) = Mass(i) < Mass(j).
	// Less(i, j int) bool
	// Swap swaps the elements with indexes i and j
	// Swap(i, j int) bool
	// Mass reports the mass/norm/value of the element indexed by i.
	Mass(i int) float64
}

// NOTE All the methods that follow assume the data has been previously
// sorted by mass.
