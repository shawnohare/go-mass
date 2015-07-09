package mass

type Pair struct {
	Element interface{}
	Mass    float64
}

// MakePair evaluates the input function f at the input x and returns
// the associated (element, mass) pair.
func MakePair(f func(interface{}) float64, x interface{}) *Pair {
	return &Pair{
		Element: x,
		Mass:    f(x),
	}
}
