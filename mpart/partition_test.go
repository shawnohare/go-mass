package mpart

import (
	"reflect"
	"testing"

	"github.com/shawnohare/go-mass"
	"github.com/stretchr/testify/assert"
)

func TestMakePartition(t *testing.T) {
	var err error
	var s mass.Slice
	// A partition variable.

	// Test a basic two cell equal width partition.
	s = mass.Slice{
		{0, 1},
		{1, 2.2},
		{2, 8.2},
		{3, 9},
	}
	partition, err := MakeMinSizeCells(s, 2, 0)
	assert.Nil(t, err)
	assert.Equal(t, []int{0, 2, 4}, partition.Indices)
	assert.Equal(t, []float64{1, 5, 9}, partition.Points)

	// Ensure that Make* and Make produce the same output.
	p2, err := Make(s, partition.Indices)
	assert.Equal(t, partition.Indices, p2.Indices)
	assert.True(t, reflect.DeepEqual(partition.Cells, p2.Cells))

	// Test the resulting cells the actual partition from the definition.
	expectedPar := []mass.Slice{
		{{0, 1}, {1, 2.2}},
		{{2, 8.2}, {3, 9}},
	}
	assert.True(t, reflect.DeepEqual(partition.Cells, expectedPar))
}

func TestMinSizeCells(t *testing.T) {
	var err error
	var s mass.Slice
	// A partition variable.

	// Test that partitioning the empty mass.Slice returns an error.
	_, err = DefineMinSizeCells(mass.Slice{}, 1, 0)
	assert.NotNil(t, err)

	// Test that the cells can be appropriately re-sized.
	s = mass.Slice{
		{nil, 1},
		{nil, 2},
		{nil, 2},
		{nil, 8},
		{nil, 9},
		{nil, 10},
	}
	p, err := DefineMinSizeCells(s, 3, 1)
	assert.Nil(t, err)
	assert.Equal(t, p.Indices, []int{0, 3, 4, 6})
}
