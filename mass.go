// Package mass provides interfaces and function to perform
// set and statistical operations on data that has some notion of mass
//  Examples of such data include:
//
// - Planets and their masses.
// - Functions and their norms.
// - A collection of random variables and an associated response variable.
// - A key, numeric value pair.
//
// A typical structure implementing the mass.Collection interface is a slice
// []A, where A is a type with some attribute corresponding to a real number.
// This attribute then induces a total ordering on []A that we might wish to:
//
// - Sort by.
// - Partition by.
// - Select subsets by.
package mass

// A Collection is a type that can sorted/partitioned/fetched from and
// represents an ordered set of objects that can be mapped to a real number.
type Collection interface {
	// Len is the number of elements in a the collection
	Len() int
	// Get the ith element.
	Get(i int) interface{}
	// Get the ith element's mass.
	Mass(i int) float64
}
