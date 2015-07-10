package rel

import (
	"math/rand"

	
)

// Sample n := size elements from a list.  If replace is true,
// then sample with replacement.
func Sample(c List, size int, replace bool) Slice {
	if size > c.Len() {
		size = c.Len()
	}

	var sample Slice
	if replace {
		sample = sampleWithReplacement(c, size)
	} else {
		sample = sampleWithoutReplacement(c, size)
	}

	return sample
}

// sample without replacement using a permutation.  Good
// for small list lengths.
func sampleWithReplacement(c List, size int) Slice {
	sample := make(Slice, size)
	for i := range sample {
		j := rand.Intn(c.Len())
		sample[i] = &Pair{
			c.Get(j),
			c.Val(j),
		}
	}

	return sample
}

// This can be slow for large list sizes.
func sampleWithoutReplacement(c List, size int) Slice {
	sample := make(Slice, size)
	perm := rand.Perm(c.Len())[:size]
	for i, j := range perm {
		sample[i] = &Pair{
			c.Get(j),
			c.Val(j),
		}
	}

	return sample
}
