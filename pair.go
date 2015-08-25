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

// MakePair constructs a new Pair instance with initialized with the
// provided values.  Currently it provides the same functionality as NewPair.
func MakePair(elem interface{}, val float64) *Pair {
	return &Pair{
		Element: elem,
		Value:   val,
	}
}
