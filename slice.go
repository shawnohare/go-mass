package rel

// A Slice is a concrete implementation of a List interface
// that still affords standard slice mechanics.
type Slice []*Pair

// Implement a List and sort interface.
func (s Slice) Get(i int) interface{}                 { return s[i].Element }
func (s Slice) Set(i int, e interface{}, val float64) { s[i] = NewPair(e, val) }
func (s Slice) Val(i int) float64                     { return s[i].Value }
func (s Slice) Len() int                              { return len(s) }
func (s Slice) Less(i, j int) bool                    { return s.Val(i) < s.Val(j) }
func (s Slice) Swap(i, j int)                         { s[i], s[j] = s[j], s[i] }

// MakeSlice will produce a Slice instance either from an input List
// or []float64 slice.  If the input is type []float64, then the
// underlying Element is an empty struct.
func MakeSlice(l interface{}) Slice {

	var s Slice

	switch t := l.(type) {
	case List:
		s = makeSliceFromList(t)
	case []float64:
		s = makeSliceFromFloat64s(t)
	}

	return s
}

func makeSliceFromList(l List) Slice {
	// Create a new Slice struct that represents the input List.
	s := make(Slice, l.Len())

	// Populate the list struct with the input List's data.
	for i := range s {
		s.Set(i, l.Get(i), l.Val(i))
	}

	return s
}

func makeSliceFromFloat64s(l []float64) Slice {
	s := make(Slice, len(l))
	for i := range l {
		s.Set(i, struct{}{}, l[i])
	}

	return s
}
