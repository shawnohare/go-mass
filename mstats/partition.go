package mstats

import (
	"errors"
	"github.com/shawnohare/go-mass"
	"math"
)

// A Partition of a mass.Collection interface into cells.
// We always view the collection, via its natural embedding, as a discrete
// subset of the reals.  All partition definitions are actually a family
// of sub-interval cell partitions that induce the same discrete partition
// over the collection.  That is, individual partition functions
// return a list of slice points for the collection that defines
// the partition.
type Partition struct {
	// Slice points that define the partition.
	Indices []int
	// Points on the Real line that induce the partition.
	Points []float64
	// The actual partition of the collection into discrete cells.
	Cells []mass.Slice
}

// Make produces a slice of sub-collections according
// to the input partition indices.  This function assumes that the
// collection has been previously sorted.
func MakePartition(c mass.Collection, partition []int) (*Partition, error) {
	var err error

	cells, err := makeCells(c, partition)
	points := makePoints(c, partition)

	p := &Partition{
		Indices: partition,
		Points:  points,
		Cells:   cells,
	}

	return p, err
}

func makeCells(c mass.Collection, partition []int) ([]mass.Slice, error) {
	if len(partition) < 2 || c.Len() < len(partition)-1 {
		err := errors.New("Could not make cells.")
		return nil, err
	}

	s := mass.MakeSlice(c)
	numCells := len(partition) - 1
	ss := make([]mass.Slice, numCells)
	for i := 0; i < numCells; i++ {
		start, end := partition[i], partition[i+1]
		ss[i] = s[start:end]
	}

	return ss, nil
}

// makePoints a slice-index partition into
// a sub-interval partition whose cells have maximum width.
func makePoints(c mass.Collection, par []int) []float64 {
	points := make([]float64, len(par))
	for i, idx := range par {
		if idx == c.Len() {
			idx = idx - 1
		}
		points[i] = c.Mass(idx)
	}

	return points
}

// MakeEqualWidthCells produces a Partition instance with populated cells
// defined by the EqualWidthCells partition.
func MakeEqualWidthCells(c mass.Collection, n int) (*Partition, error) {
	return MakeMinSizeCells(c, n, 0)
}

// EqualWidthCells partitions the collection into equal width subintervals.
// It is the partial evaluation of MinSizedCells.
func DefineEqualWidthCells(c mass.Collection, n int) (*Partition, error) {
	return DefineMinSizeCells(c, n, 0)
}

// MakeMinSizedCells calls MinSizedCells to obtain a partition
// definition and then makes the specified partition.
func MakeMinSizeCells(c mass.Collection, n, k int) (*Partition, error) {
	p, err := DefineMinSizeCells(c, n, k)
	if err != nil {
		return nil, err
	}
	cells, err := makeCells(c, p.Indices)
	p.Cells = cells
	return p, err
}

// MinSizeCells produces a partition of the previously sorted
// input mass.Collection into n subinterval cells each of which has at least
// k elements. The partition tries to be as close to an equal width
// partition as possible. In particular, when k = 0, the result is
// an equal width partition.
func DefineMinSizeCells(c mass.Collection, n, k int) (*Partition, error) {
	N := c.Len()
	if n < 1 || k < 0 || N < n*k || N < 2 || c.Mass(N-1) == c.Mass(0) {
		err := errors.New("mass.Collection cannot be partitioned.")
		return nil, err
	}

	originalWidth := (c.Mass(N-1) - c.Mass(0)) / float64(n)
	// Define the index i of the current element we must assign to some cell.
	i := 0
	// The partition consists of n + 1 elements so that a pairwise traversal
	// yields approrpiate slice points to define cells.
	points := make([]float64, n+1)
	points[0] = c.Mass(0)
	points[n] = c.Mass(N - 1)
	parIdxs := make([]int, n+1)
	parIdxs[0] = 0
	parIdxs[n] = N
	for j := 1; j < n; j++ {
		// Determine the right cell edge for an initially equal width partition.
		// FIXME delete commented out code eventually when things stabilize.
		// originalEdge := c.Mass(i) + (c.Mass(N-1) - c.Mass(i)/float64(n-j))
		originalEdge := c.Mass(0) + float64(j)*originalWidth
		i += k
		// Set the index of element to use for determining right bin edge.
		// We use i - 1 so that the ith element and (i-1)th element belong
		// to the same cell if their masses are equal.
		i2 := i - 1
		if i2 < 0 {
			i2 = 0
		}
		rightEdge := math.Max(originalEdge, c.Mass(i2))
		points[j] = rightEdge

		// Add elements into the cell until none belong or there are too
		// few remaining elements to distribute amongst the remaining cells.
		elemsLeft := N - i
		elemsToDistribute := k * (n - j - 1)
		for c.Mass(i) <= rightEdge && elemsLeft > elemsToDistribute {
			i++
			elemsLeft--
		}
		parIdxs[j] = i
	}

	def := &Partition{
		Indices: parIdxs,
		Points:  points,
	}

	return def, nil
}
