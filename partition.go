package mass

// Bin is an arbitrary interval of masses.  It specifies the
// index range of data whose mass falls in the bin as well as
// mass endpoints.  Whether the bin is open, closed, or clopen
// depends on context and is not encoded explicitly.
type Bin struct {
	MinEdge  float64 // left endpoint of interval
	MaxEdge  float64 // right endpoint of interval
	MinIndex int     // index of initial element in bin
	MaxIndex int     // index of final element in bin
}

// EqualWidthBins provides the index ranges of
// the sorted (by mass) data that fall into n := numBins
// equal width sub-intervals over the range of the data.
// The output is a slice of n Bin objects that encode
func EqualWidthBins(data Interface, numBins int) []Bin {
	var bins []Bin
	// TODO write
	return bins
}
