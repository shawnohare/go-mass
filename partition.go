package mass

import (
	"fmt"
)

// Bin is an arbitrary interval of masses.  It specifies the
// index range of data whose mass falls in the bin as well as
// mass endpoints.  Whether the bin is open, closed, or clopen
// depends on context and is not encoded explicitly.
type Bin struct {
	MinEdge  float64 // left endpoint of interval
	MaxEdge  float64 // right endpoint of interval
	Width    float64 // length of subinterval defining the bin
	MinIndex int     // index of initial element in bin
	MaxIndex int     // index of final element in bin
	Type     string  // whether bin is (), [], (], [)
}

// EqualWidthBins provides the index ranges of
// the sorted (by mass) data that fall into n := numBins
// equal width sub-intervals over the range of the data.
// The output is a slice of n Bin objects that encode
func EqualWidthBins(data Interface, numBins int) ([]Bin, error) {
	var err error
	var bins []Bin
	l := data.Len()

	switch {
	case numBins < 1:
		err = fmt.Errorf("Number of bins must be > 0.")
	case l < 2:
		err = fmt.Errorf("Too few data points to partition.  data.Len()=%d < 2.", l)
	case data.Mass(0) < data.Mass(l-1):
		err = fmt.Errorf("Data not sorted.  Maximal mass %f < minimal mass %f.", data.Mass(0), data.Mass(l-1))
	case data.Mass(0) == data.Mass(l-1):
		err = fmt.Errorf("Cannot create partition. Minimal mass %f = maximal mass %f.", data.Mass(0), data.Mass(l-1))
	default:
		min := data.Mass(0)
		max := data.Mass(l - 1)
		bins = make([]Bin, numBins)
		width := (max - min) / float64(numBins)
		rightIndex := 0

		for i := 0; i < numBins; i++ {
			left := min + float64(i)*width
			right := left + width
			leftIndex := rightIndex

			if i < numBins-1 {
				for data.Mass(rightIndex) < right {
					rightIndex++
				}
				bin := Bin{
					MinEdge:  left,
					MaxEdge:  right,
					Width:    width,
					MinIndex: leftIndex,
					MaxIndex: rightIndex,
					Type:     "[ )",
				}
				bins[i] = bin
			} else {
				// Final bin
				bin := Bin{
					MinEdge:  left,
					MaxEdge:  right,
					Width:    width,
					MinIndex: leftIndex,
					MaxIndex: l - 1,
					Type:     "[  ]",
				}
				bins[i] = bin
			}
		}
	}
	return bins, err
}
