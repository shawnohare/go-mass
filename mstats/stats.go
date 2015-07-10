// Package mstats provides some basic statistical functions that
// act over mass.mass.Collection interfaces to compute statistics on
// datasets with respect to their mass. Values are assumed
// to be in the extended real numbers (which includes +, - infinty).
// When necessary, pre-sorting is assumed.
package mstats

import (
	"math"

	"github.com/shawnohare/go-mass"
)

// ComputeStat computes a real-valued statistic over a mass.Collection.
func ComputeStat(c mass.Collection, stat func([]float64) float64) float64 {
	xs := make([]float64, c.Len())
	for i := range xs {
		xs[i] = c.Mass(i)
	}

	return stat(xs)
}

// Mean returns the mean mass.
func Mean(data mass.Collection) float64 {
	var totalMass float64
	var mean float64

	l := data.Len()

	if l == 0 {
		return math.NaN()
	}

	for i := 0; i < l; i++ {
		totalMass += data.Mass(i)
	}
	mean = totalMass / float64(l)
	return mean
}

// Min returns the element with the minimal mass.
func Min(data mass.Collection) interface{} {
	// For finite sets the minimal value is equal to the
	// greatest lower bound, or infimum. The infimum for
	// the empty set is vacuously positive infinity.
	if data.Len() == 0 {
		return nil
	}

	min := data.Mass(0)
	index := 0
	for i := 1; i < data.Len(); i++ {
		mass := data.Mass(i)
		if mass < min {
			index = i
			min = mass
		}
	}

	return data.Get(index)
}

// Max returns the element with the minimal mass.
func Max(data mass.Collection) interface{} {
	// For finite sets the maximal value is equal to the
	// least upper bound, or supremum.  The supremum for
	// the empty set is vacuously negative infinity.
	if data.Len() == 0 {
		return nil
	}

	max := data.Mass(0)
	index := 0
	for i := 1; i < data.Len(); i++ {
		mass := data.Mass(i)
		if mass > max {
			index = i
			max = mass
		}
	}

	return data.Get(index)
}
