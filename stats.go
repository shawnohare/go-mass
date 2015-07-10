package rel

import (
	"math"
)

// ComputeStat computes a real-valued statistic over a List.
func ComputeStat(c List, stat func([]float64) float64) float64 {
	xs := make([]float64, c.Len())
	for i := range xs {
		xs[i] = c.Val(i)
	}

	return stat(xs)
}

// Mean returns the mean
func Mean(data List) float64 {
	var totalMass float64
	var mean float64

	l := data.Len()

	if l == 0 {
		return math.NaN()
	}

	for i := 0; i < l; i++ {
		totalMass += data.Val(i)
	}
	mean = totalMass / float64(l)
	return mean
}

// Min returns the element with the minimal
func Min(data List) interface{} {
	// For finite sets the minimal value is equal to the
	// greatest lower bound, or infimum. The infimum for
	// the empty set is vacuously positive infinity.
	if data.Len() == 0 {
		return nil
	}

	min := data.Val(0)
	index := 0
	for i := 1; i < data.Len(); i++ {
		mass := data.Val(i)
		if mass < min {
			index = i
			min = mass
		}
	}

	return data.Get(index)
}

// Max returns the element with the minimal
func Max(data List) interface{} {
	// For finite sets the maximal value is equal to the
	// least upper bound, or supremum.  The supremum for
	// the empty set is vacuously negative infinity.
	if data.Len() == 0 {
		return nil
	}

	max := data.Val(0)
	index := 0
	for i := 1; i < data.Len(); i++ {
		mass := data.Val(i)
		if mass > max {
			index = i
			max = mass
		}
	}

	return data.Get(index)
}
