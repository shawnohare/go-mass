// Statistical operations with respect the mass.  Values are assumed
// to be in the extended real numbers (which includes +, - infinty).
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

// Min returns the minimal mass from a not necessarily sorted collection.
func Min(data Interface) float64 {
	// For finite sets the minimal value is equal to the
	// greatest lower bound, or infimum. The infimum for
	// the empty set is vacuously positive infinity.
	min := math.Inf(1)
	if data.Len() > 0 {
		min = data.Mass(0)
		for i := 1; i < data.Len(); i++ {
			min = math.Min(min, data.Mass(i))
		}
	}

	return min
}
