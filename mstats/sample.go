package mstats

import (
	"math/rand"

	"github.com/shawnohare/go-mass"
)

// Sample n := size elements from a collection.  If replace is true,
// then sample with replacement.
func Sample(c mass.Collection, size int, replace bool) mass.Slice {
	if size > c.Len() {
		size = c.Len()
	}

	var sample mass.Slice
	if replace {
		sample = sampleWithReplacement(c, size)
	} else {
		sample = sampleWithoutReplacement(c, size)
	}

	return sample
}

// sample without replacement using a permutation.  Good
// for small collection lengths.
func sampleWithReplacement(c mass.Collection, size int) mass.Slice {
	sample := make(mass.Slice, size)
	for i := range sample {
		j := rand.Intn(c.Len())
		sample[i] = &mass.Pair{
			c.Get(j),
			c.Mass(j),
		}
	}

	return sample
}

// This can be slow for large collection sizes.
func sampleWithoutReplacement(c mass.Collection, size int) mass.Slice {
	sample := make(mass.Slice, size)
	perm := rand.Perm(c.Len())[:size]
	for i, j := range perm {
		sample[i] = &mass.Pair{
			c.Get(j),
			c.Mass(j),
		}
	}

	return sample
}
