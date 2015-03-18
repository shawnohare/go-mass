// Statistical operations with respect the mass.  Values are assumed
// to be in the extended real numbers (which includes +, - infinty).
// When necessary, pre-sorting is assumed.
package mass

import (
	"math"
)

// Mean returns the mean mass.
func Mean(data Interface) float64 {
	var totalMass float64
	var mean float64
	l := data.Len()

	for i := 0; i < l; i++ {
		totalMass += data.Mass(i)
	}

	// Deal with empty set case.
	if l == 0 {
		mean = math.NaN()
	} else {
		mean = totalMass / l
	}

	return mean
}

// Min returns the minimal mass and corresponding index.
func Min(data Interface) (float64, int) {
	// For finite sets the minimal value is equal to the
	// greatest lower bound, or infimum. The infimum for
	// the empty set is vacuously positive infinity.
	min := math.Inf(1)
	index := -1
	if data.Len() > 0 {
		min = data.Mass(0)
		index = 0
		for i := 1; i < data.Len(); i++ {
			mass := data.Mass(i)
			if mass < min {
				index = i
				min = mass
			}
		}
	}

	return min, index
}

// Max returns the maximal mass and corresponding index.
func Max(data Interface) (float64, int) {
	// For finite sets the maximal value is equal to the
	// least upper bound, or supremum.  The supremum for
	// the empty set is vacuously negative infinity.
	max := math.Inf(-1)
	index := -1
	if data.Len() > 0 {
		max := data.Mass(0)
		index = 0
		for i := 1; i < data.Len(); i++ {
			mass := data.Mass(i)
			if mass > max {
				index = i
				max = mass
			}
		}
	}

	return max, index
}
