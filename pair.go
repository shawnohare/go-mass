package rel

// A Pair consisting of an element and its real value.
type Pair struct {
	Element interface{}
	Value   float64
}

// MakePair evaluates the input function f at the input x and returns
// the associated (element, mass) pair.
func MakePair(f func(interface{}) float64, x interface{}) *Pair {
	return &Pair{
		Element: x,
		Value:   f(x),
	}
}
