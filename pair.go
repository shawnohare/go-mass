package rel

// A Pair consisting of an element and its real value.
type Pair struct {
	Element interface{}
	Value   float64
}

// NewPair creates a *Pair from the input interface and value.
func NewPair(elem interface{}, val float64) *Pair {
	return &Pair{
		Element: elem,
		Value:   val,
	}
}
